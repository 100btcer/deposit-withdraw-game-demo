import { Modal } from "antd";
import { useEffect, useMemo, useState, useRef } from "react";
import type { FC } from "react";
import { useTranslation } from "react-i18next";
import { parseUnits, formatUnits, MaxUint256 } from "ethers";
import styles from "./index.module.scss";
import closeIcon from "../../images/close-icon.png";
import tabBottom from "../../images/tab-bottom.png";
import MaxText from "../../images/MAX-text.png";
import "../../App.css";
import {
  useRequestUserInfo,
  useRequestDepositeList,
  useRequestGetValideNumber,
} from "../../apis/hooks";
import { getParsedEthersError } from "@enzoferey/ethers-error-parser";
import RangersCoin from "../../images/rangers-coin.png";
import veLP from "../../images/tiket-coin.png";
import browserIcon from "../../images/goBrowser.png";
import cls from "classnames";
import { useContractWrite, useContractRead, useAccount, erc20ABI } from "wagmi";
import TipModal from "../TipModal";
import { scanUrl, mmAddress, lpAddress } from "../../conifg/scan";
import mmAbi from "../../conifg/mmabi.json";
import { network, useChangeNetwork } from "../Content";
import withdrawCn from "../../../src/images/langImages/withdraw-cn.png";
import withdrawEn from "../../../src/images/langImages/withdraw-en.png";
import withdrawKr from "../../../src/images/langImages/withdraw-kr.png";
import stakeCn from "../../../src/images/langImages/stake-cn.png";
import stakeEn from "../../../src/images/langImages/stake-en.png";
import stakeKr from "../../../src/images/langImages/stake-kr.png";
import splitLine from "../../../src/images/split-line.png";
import { Empty } from "../BagModal";
const drawTextMap: Record<string, string> = {
  kr: withdrawKr,
  en: withdrawEn,
  cn: withdrawCn,
};

const stakeMap: Record<string, string> = {
  kr: stakeKr,
  en: stakeEn,
  cn: stakeCn,
};

enum Active {
  Deposite = 1,
  Back = 2,
}

enum DayPick {
  Seven = 7,
  fourteen = 14,
  thirty = 30,
  ninety = 90,
  eityO = 180,
  oneYear = 365,
}

const RecountMap = {
  [DayPick.Seven]: 1.2,
  [DayPick.fourteen]: 1.39,
  [DayPick.thirty]: 1.6,
  [DayPick.ninety]: 2.11,
  [DayPick.eityO]: 1.75,
  [DayPick.oneYear]: 1.16,
};

const RealRecountMap = {
  [DayPick.Seven]: 0.000989088,
  [DayPick.fourteen]: 0.001149292,
  [DayPick.thirty]: 0.001323427,
  [DayPick.ninety]: 0.001741351,
  [DayPick.eityO]: 0.001448804,
  [DayPick.oneYear]: 0.000961226,
};

const RecountIDMap = {
  [DayPick.Seven]: "1",
  [DayPick.fourteen]: "2",
  [DayPick.thirty]: "3",
  [DayPick.ninety]: "4",
  [DayPick.eityO]: "5",
  [DayPick.oneYear]: "6",
};

const TabOne = ({
  handlePlayOtherButtonSound,
  lpNumber,
  lpNumberRefecth,
  TiketNumberRefetch,
}: any) => {
  const [activeDay, setActiveDay] = useState<DayPick>(DayPick.Seven);
  const { t, i18n } = useTranslation("translation");
  const lang = i18n.language;
  const firstRender = useRef(true);
  const [value, setvalue] = useState<any>(lpNumber);
  const { address } = useAccount();
  const [open, setOpen] = useState<boolean>(false);
  const { handleChangeNetwork } = useChangeNetwork();
  const handleInputChange = (event: React.FormEvent<HTMLInputElement>) => {
    const a = event.currentTarget.value.replace(/^0{1,}/g, "");
    if (+a > lpNumber) {
      setvalue(lpNumber);
    } else {
      setvalue(a);
    }
  };
  const { data, isLoading, isSuccess, writeAsync, error } = useContractWrite({
    address: mmAddress,
    abi: mmAbi,
    chainId: network,
    functionName: "deposit",
    args: [
      RecountIDMap[activeDay].toString(),
      value ? parseUnits(value.toString(), 18) : "0x",
    ],
  });

  const { data: AllowenceData, refetch: allowanceRefetch } = useContractRead({
    address: lpAddress,
    abi: erc20ABI,
    chainId: network,
    functionName: "allowance",
    args: [address || "0x", mmAddress],
  });
  if (AllowenceData) {
    console.log(
      formatUnits(AllowenceData?.toString(), 18),
      value,
      "ppwww",
      +formatUnits(AllowenceData?.toString(), 18) >= +value
    );
  }
  const {
    isLoading: approveloading,
    isSuccess: approveSuccess,
    writeAsync: approveWrite,
    error: approveError,
  } = useContractWrite({
    address: lpAddress,
    abi: erc20ABI,
    chainId: network,
    functionName: "approve",
    args: [mmAddress, MaxUint256],
  });

  const handelDeposite = async () => {
    await handleChangeNetwork();
    if (isLoading) return;
    handlePlayOtherButtonSound();
    if (
      AllowenceData &&
      +formatUnits(AllowenceData?.toString(), 18) >= +value
    ) {
      if (value) {
        setOpen(true);
        await writeAsync();
      }
    } else {
      approveWrite();
    }
  };
  useEffect(() => {
    if (approveSuccess) {
      setTimeout(() => {
        allowanceRefetch();
      }, 500);
    }
  }, [approveSuccess]);
  useEffect(() => {
    console.log(AllowenceData, "AllowenceData----");
    if (AllowenceData) {
      if (firstRender.current) {
        firstRender.current = false;
      } else {
        handelDeposite();
      }
    }
  }, [AllowenceData]);
  useEffect(() => {
    if (isSuccess) {
      setOpen(true);
      lpNumberRefecth();
      TiketNumberRefetch();
      setvalue(0);
      allowanceRefetch();
    }
  }, [isSuccess, lpNumberRefecth, TiketNumberRefetch, allowanceRefetch]);
  const handelMax = () => {
    handlePlayOtherButtonSound();
    setvalue(lpNumber);
  };
  const tipText = useMemo(() => {
    const tipError = error;
    if (isLoading) return t("transation-pending");
    if (isSuccess) return t("transtion-success");
    if (tipError && tipError.message.includes("User rejected the request"))
      return t("user-cancel-transaction");
    return t("trasaction-fail");
  }, [isLoading, isSuccess, approveloading]);
  return (
    <div className={styles["tab-one"]}>
      <TipModal
        open={open}
        handleClose={() => setOpen(false)}
        tipType={isSuccess || isLoading ? "success" : "error"}
        scanAddress={data?.hash as unknown as string}
        tipText={tipText}
      ></TipModal>
      <div className={styles["input-area"]}>
        <div className={styles["input-left"]}>
          <input type="number" value={value} onChange={handleInputChange} />
        </div>
        <div className={styles["input-right"]} onClick={handelMax}>
          <img src={MaxText} alt="" />
        </div>
      </div>
      <div className={styles["input-tip"]}>
        {t("can-stake")}
        <img src={RangersCoin} alt="" />
        <span>{lpNumber}</span>
      </div>
      <img src={splitLine} alt="" className="split-line" />
      <div className={styles["time-chose"]}>
        <div className={styles["chose-title"]}>{t("chose-stake-time")}</div>
        <div className={styles["time-box"]}>
          <div
            className={cls(
              styles["time-box-item"],
              activeDay === DayPick.Seven && styles["time-box-item-active"]
            )}
            onClick={() => setActiveDay(DayPick.Seven)}
          >
            7 {t("days")}
          </div>
          <div
            className={cls(
              styles["time-box-item"],
              activeDay === DayPick.fourteen && styles["time-box-item-active"]
            )}
            onClick={() => setActiveDay(DayPick.fourteen)}
          >
            14 {t("days")}
          </div>
          <div
            className={cls(
              styles["time-box-item"],
              activeDay === DayPick.thirty && styles["time-box-item-active"]
            )}
            onClick={() => setActiveDay(DayPick.thirty)}
          >
            30 {t("days")}
          </div>
          <div
            className={cls(
              styles["time-box-item"],
              activeDay === DayPick.ninety && styles["time-box-item-active"]
            )}
            onClick={() => setActiveDay(DayPick.ninety)}
          >
            90 {t("days")}
          </div>
          <div
            className={cls(
              styles["time-box-item"],
              activeDay === DayPick.eityO && styles["time-box-item-active"]
            )}
            onClick={() => setActiveDay(DayPick.eityO)}
          >
            180 {t("days")}
          </div>
          <div
            className={cls(
              styles["time-box-item"],
              activeDay === DayPick.oneYear && styles["time-box-item-active"]
            )}
            onClick={() => setActiveDay(DayPick.oneYear)}
          >
            365 {t("days")}
          </div>
        </div>
        <div className={styles["time-chose-tip"]}>
          {t("time-diffrent-call")}
        </div>
      </div>
      <img src={splitLine} alt="" className="split-line" />
      <div className={styles["deposite-tip"]}>
        {t("predicr-get-tiket")}:
        <img src={veLP} alt="" />
        <div className={styles["predict-ticket"]}>
          {value
            ? parseFloat(
                (value * RealRecountMap[activeDay] * activeDay).toFixed(2)
              )
            : 0}
        </div>
      </div>
      <div className={styles["deposite-button"]} onClick={handelDeposite}>
        <img
          src={stakeMap[lang] || stakeEn}
          alt=""
          className={cls(
            styles["stake-image-text"],
            lang === "cn" && styles["cn-stake"],
            lang === "kr" && styles["kr-stake"]
          )}
        />
      </div>
    </div>
  );
};

const TabTwo = ({ handlePlayOtherButtonSound }: any) => {
  const { t, i18n } = useTranslation("translation");
  const lang = i18n.language;
  const [open, setOpen] = useState<boolean>(false);
  const { handleChangeNetwork } = useChangeNetwork();
  const { data: dataAmount, refetch } = useRequestGetValideNumber();
  const {
    data: contractData,
    isLoading,
    isSuccess,
    writeAsync,
    error,
  } = useContractWrite({
    address: mmAddress,
    abi: mmAbi,
    chainId: network,
    functionName: "withdraw",
  });
  const handleGetBack = async () => {
    await handleChangeNetwork();
    if (isLoading) return;
    handlePlayOtherButtonSound();
    writeAsync();
    setOpen(true);
  };
  useEffect(() => {
    if (isSuccess) {
      refetch();
    }
  }, [isSuccess]);
  const tipText = useMemo(() => {
    if (isLoading) return t("transation-pending");
    if (isSuccess) return t("transtion-success");
    if (error && error.message.includes("User rejected the request")) {
      return t("user-cancel-transaction");
    }
    return t("trasaction-fail");
  }, [isLoading, isSuccess]);
  const { data } = useRequestUserInfo();
  const { data: data2 } = useRequestDepositeList();

  return (
    <div className={styles["tab-two"]}>
      <TipModal
        open={open}
        handleClose={() => setOpen(false)}
        tipType={isSuccess || isLoading ? "success" : "error"}
        scanAddress={contractData?.hash as unknown as string}
        tipText={tipText}
      ></TipModal>
      <div className={styles["top-input-tip"]}>{t("can-back")}</div>
      <div className={styles["input-area"]}>
        <div className={styles["input-left"]}>
          <input
            type="number"
            pattern="number"
            disabled
            value={dataAmount?.amount || 0}
          />
        </div>
        <div className={styles["input-right"]} onClick={handleGetBack}>
          <img
            src={drawTextMap[lang] || withdrawEn}
            alt=""
            className={styles["withdraw-text"]}
          />
        </div>
      </div>
      <div className={styles["input-tip"]}>
        {t("is-stake")}
        <img src={RangersCoin} alt="" />
        <span>{data?.deposit_amount || "--"}</span>
      </div>
      <img src={splitLine} alt="" className="split-line" />
      <div className={styles["deposite-reocrd-title"]}>
        {t("deposite-record-title")}
      </div>
      <div className={styles["list-title"]}>
        <div>{t("deposite-time")}</div>
        <div>{t("amount")}</div>
        <div>{t("expire-time")}</div>
        <div>{t("browser")}</div>
      </div>
      <div className={styles["list-content"]}>
        {data2?.list ? (
          data2?.list.map((item, index) => (
            <div className={styles["list-content-item"]} key={index}>
              <div>{item.create_time}</div>
              <div className={styles["amount-number"]}>
                {item.deposit_amount}
              </div>
              <div>{item.lock_day}</div>
              <div>
                <a
                  target="_blank"
                  rel="noopener noreferrer"
                  href={`${scanUrl}${item.hash}`}
                >
                  <img src={browserIcon} alt="" />
                </a>
              </div>
            </div>
          ))
        ) : (
          <Empty />
        )}
      </div>
    </div>
  );
};

interface ModalProsps {
  open: boolean;
  handleClose: () => void;
  handlePlayOtherButtonSound: () => void;
  lpNumber: number;
  lpNumberRefecth: () => void;
  TiketNumberRefetch: () => void;
}

const Index: FC<ModalProsps> = ({
  open,
  handleClose,
  handlePlayOtherButtonSound,
  lpNumber,
  lpNumberRefecth,
  TiketNumberRefetch,
}) => {
  const [active, setActive] = useState<Active>(Active.Deposite);
  const { t } = useTranslation("translation");
  return (
    <Modal
      wrapClassName={styles["body-wrrape"]}
      open={open}
      closable={false}
      footer={null}
      width={368}
    >
      <img
        src={closeIcon}
        alt=""
        className={styles["close-icon"]}
        onClick={handleClose}
      />
      <div className={styles["title-area"]}>
        <div className={styles["title-text"]}>{t("stake/back")}</div>
      </div>
      <div className={styles["tab-area"]}>
        <div className={styles["tab-card"]}>
          <div
            className={cls(
              styles["card-item"],
              active === Active.Deposite && styles["active-item"]
            )}
            onClick={() => setActive(Active.Deposite)}
          >
            {t("deposite")}
          </div>
          <div
            className={cls(
              styles["card-item"],
              active === Active.Back && styles["active-item"]
            )}
            onClick={() => setActive(Active.Back)}
          >
            {t("back")}
          </div>
        </div>
        <div className={styles["tab-area-bottom"]}>
          <img src={tabBottom} alt="" />
        </div>
        {active === Active.Deposite && (
          <TabOne
            handlePlayOtherButtonSound={handlePlayOtherButtonSound}
            lpNumber={lpNumber}
            lpNumberRefecth={lpNumberRefecth}
            TiketNumberRefetch={TiketNumberRefetch}
          ></TabOne>
        )}
        {active === Active.Back && (
          <TabTwo
            handlePlayOtherButtonSound={handlePlayOtherButtonSound}
          ></TabTwo>
        )}
      </div>
    </Modal>
  );
};

export default Index;
