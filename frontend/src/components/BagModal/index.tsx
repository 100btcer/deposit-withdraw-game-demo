import { Spin, Modal, message } from "antd";
import { useState, useMemo } from "react";
import TipModal from "../TipModal";
import type { FC } from "react";
import cls from "classnames";
import { useTranslation } from "react-i18next";
import styles from "./index.module.scss";
import closeIcon from "../../images/close-icon.png";
import tabBottom from "../../images/tab-bottom.png";
import { useRequestRewardList } from "../../apis/hooks";
import Mix from "../../images/Mix.png";
import { handleFormatList } from "../Content";
import { useMutationGetLottery } from "../../apis/hooks";
import DrawCn from "../../../src/images/langImages/claim-cn.png";
import DrawEn from "../../../src/images/langImages/claim-en.png";
import DrawKr from "../../../src/images/langImages/claim-kr.png";
import empytyIcon from "../../../src/images/empytyIcon.png";

const drawTextMap: Record<string, string> = {
  kr: DrawKr,
  en: DrawEn,
  cn: DrawCn,
};

enum Active {
  TOKEN = 1,
  NFT = 2,
  Record = 3,
}

export const Empty = () => {
  const { t } = useTranslation();
  return (
    <div className={styles["empty-container"]}>
      <img src={empytyIcon} alt="" className={styles["empyty-icon"]} />
      {t("empty")}
    </div>
  );
};

interface ModalProsps {
  open: boolean;
  handleClose: () => void;
  handlePlayOtherButtonSound: () => void;
}

const MySpinning = ({ spinning }: { spinning: boolean }) => {
  return (
    <div className={styles["my-spining"]}>
      <Spin spinning={spinning}></Spin>
    </div>
  );
};

const CardItem = ({
  name = "MIX",
  number = 5000,
  img = { Mix },
  id,
  type,
  link,
  token_id,
  handleGet,
}: any) => {
  const { t, i18n } = useTranslation();

  const lang = i18n.language;
  const showName = useMemo(() => {
    if (+token_id === 7) {
      return t("egg-normal");
    }
    if (+token_id === 8) {
      return t("egg-privte");
    }
    return name;
  }, [name, token_id, t]);
  const content = (
    <div className={styles["card-item-wrape"]}>
      <div className={styles["card-item-p"]}>
        <img src={img} alt="" className={styles["card-item-p-image"]} />
        <div className={styles["name-area"]}>{showName}</div>
        <div className={styles["count-area"]}>
          {type === 1 ? number : `X${number}`}
        </div>
      </div>
      {type === 1 && (
        <div className={styles["button-area"]} onClick={() => handleGet(id)}>
          <img
            src={drawTextMap[lang] || DrawEn}
            alt=""
            className={styles["claim-text"]}
          />
        </div>
      )}
    </div>
  );
  if (link) {
    return (
      <a href={link} target="_blank" rel="noopener noreferrer">
        {content}
      </a>
    );
  }
  return content;
};

const TabRecord = ({
  list,
  isLoading,
}: {
  list: Array<any>;
  isLoading: boolean;
}) => {
  const { t } = useTranslation();
  const getName = (token_id: number, name: string) => {
    if (token_id === 7) {
      return t("egg-normal");
    }
    if (token_id === 8) {
      return t("egg-privte");
    }
    return name;
  };
  return (
    <div className={styles["tabRecord"]}>
      <div className={styles["list-title"]}>
        <div>{t("time")}</div>
        <div>{t("reward")}</div>
        <div>{t("amount")}</div>
      </div>
      <div className={styles["list-content"]}>
        {isLoading ? (
          <MySpinning spinning={isLoading}></MySpinning>
        ) : list && list.length > 0 ? (
          list.map((item, index) => {
            const showName = getName(item.token_id, item.name);
            return (
              <div className={styles["list-content-item"]} key={index}>
                <div>{item.create_time}</div>
                <div>{showName}</div>
                <div className={styles["amount-number"]}>{item.amount}</div>
              </div>
            );
          })
        ) : (
          <Empty></Empty>
        )}
      </div>
    </div>
  );
};

const Index: FC<ModalProsps> = ({
  open,
  handleClose,
  handlePlayOtherButtonSound,
}) => {
  const [active, setActive] = useState<Active>(Active.TOKEN);
  const [resHash, setResHash] = useState("");
  const [Tipopen, setOpen] = useState(false);
  const { mutateAsync, isLoading: getLoading } = useMutationGetLottery();
  const handleGet = async (id: string) => {
    if (getLoading) return;
    handlePlayOtherButtonSound();
    const res = await mutateAsync({ id });
    setOpen(true);
    setResHash(res as unknown as string);
    refetch();
  };
  const paramsType = useMemo(() => {
    return active === Active.Record ? undefined : active;
  }, [active]);

  const { data, refetch, isLoading } = useRequestRewardList({
    type: paramsType,
  });
  const list1 = useMemo(() => {
    if (data?.list && active !== Active.Record) {
      return handleFormatList(data?.list);
    }
    if (data?.list) {
      return data?.list;
    }
    return [];
  }, [data]);
  const { t } = useTranslation();
  return (
    <Modal
      wrapClassName={styles["body-wrrape"]}
      open={open}
      closable={false}
      footer={null}
      width={368}
      centered
    >
      <TipModal
        tipText={`${t("go-get")}  ${t("claimed")}`}
        scanAddress={resHash}
        open={Tipopen}
        handleClose={() => setOpen(false)}
        tipType="success"
      ></TipModal>
      <img
        src={closeIcon}
        alt=""
        className={styles["close-icon"]}
        onClick={handleClose}
      />
      <div className={styles["title-area"]}>
        <div className={styles["title-text"]}>{t("my-reward")}</div>
      </div>
      <div className={styles["tab-area"]}>
        <div className={styles["tab-card"]}>
          <div
            className={cls(
              styles["card-item"],
              active === Active.TOKEN && styles["active-item"]
            )}
            onClick={() => setActive(Active.TOKEN)}
          >
            Token
          </div>
          <div
            className={cls(
              styles["card-item"],
              active === Active.NFT && styles["active-item"]
            )}
            onClick={() => setActive(Active.NFT)}
          >
            NFT
          </div>
          <div
            className={cls(
              styles["card-item"],
              active === Active.Record && styles["active-item"]
            )}
            onClick={() => setActive(Active.Record)}
          >
            {t("reward-list")}
          </div>
        </div>
        <div className={styles["tab-area-bottom"]}>
          <img src={tabBottom} alt="" />
        </div>
        {active === Active.Record ? (
          <TabRecord list={list1} isLoading={isLoading}></TabRecord>
        ) : (
          <div className={styles["rewardList"]}>
            {isLoading ? (
              <MySpinning spinning={isLoading}></MySpinning>
            ) : list1 && list1.length > 0 ? (
              list1.map((item, index) => (
                <CardItem
                  handleGet={handleGet}
                  name={item.name}
                  id={item.id}
                  key={index}
                  token_id={item.token_id}
                  img={item.icon}
                  link={item.link}
                  number={item.amount}
                  type={item.type}
                  refetch={refetch}
                  handlePlayOtherButtonSound={handlePlayOtherButtonSound}
                ></CardItem>
              ))
            ) : (
              <Empty></Empty>
            )}
          </div>
        )}
      </div>
    </Modal>
  );
};

export default Index;
