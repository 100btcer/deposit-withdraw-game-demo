import { Modal } from "antd";
import { useState, useMemo, useEffect } from "react";
import type { FC } from "react";
import { useTranslation } from "react-i18next";
import styles from "./index.module.scss";
import closeIcon from "../../images/close-icon.png";
import veLP from "../../images/tiket-coin.png";
import cls from "classnames";
import { useContractWrite, useAccount } from "wagmi";
import { parseUnits } from "ethers";
import { useChangeNetwork } from "../Content";
import "../../App.css";
import { ticketAddress } from "../../conifg/scan";
import blueArrow from "../../images/blue-arrow.png";
import maxText from "../../images/MAX-text.png";
import ticketAbi from "../../conifg/ticketabi.json";
import TipModal from "../TipModal";
import splitLine from "../../../src/images/split-line.png";

const Times = 10;

interface ModalProsps {
  open: boolean;
  ticketNumber: number;
  handleClose: () => void;
  handlePlayDrawButton: () => void;
  handleShowWaiteVideo: (value: `0x${string}`) => void;
}

const Index: FC<ModalProsps> = ({
  open,
  handleClose,
  handlePlayDrawButton,
  handleShowWaiteVideo,
  ticketNumber,
}) => {
  const { handleChangeNetwork } = useChangeNetwork();
  const { t, i18n } = useTranslation();
  const [tipOpen, setTipOpen] = useState(false);
  const [value, setvalue] = useState<any>(0);
  const count = useMemo(() => {
    if (ticketNumber) {
      return Math.floor(ticketNumber / Times);
    }
    return 0;
  }, [ticketNumber]);
  const { address } = useAccount();
  const { data, isLoading, isSuccess, writeAsync, error } = useContractWrite({
    address: ticketAddress,
    abi: ticketAbi,
    functionName: "burn",
    args: [address, value ? parseUnits((value * Times).toString(), 18) : "0x"],
  });
  const handleInputChange = (event: React.FormEvent<HTMLInputElement>) => {
    const value = event.currentTarget.value.replace(/^0{1,}/g, "");
    if (+value <= count) {
      setvalue(value);
    } else {
      setvalue(count);
    }
  };
  const handlSubmit = async () => {
    await handleChangeNetwork();
    handlePlayDrawButton();
    writeAsync();
    setTipOpen(open);
  };
  useEffect(() => {
    if (isSuccess) {
      handleShowWaiteVideo(data?.hash || "0x");
      handleClose();
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
        open={tipOpen}
        handleClose={() => setTipOpen(false)}
        tipType={isSuccess || isLoading ? "success" : "error"}
        scanAddress={data?.hash as unknown as string}
        tipText={tipText}
      ></TipModal>
      <img
        src={closeIcon}
        alt=""
        className={styles["close-icon"]}
        onClick={handleClose}
      />
      <div className={styles["title-area"]}>
        <div className={styles["title-text"]}>
          {t("cost")} <img src={veLP} alt="" /> <span>{Times}</span>/
          {t("times")}
        </div>
      </div>
      <div className={styles["time-input"]}>
        <img
          src={blueArrow}
          onClick={() => {
            if (+value > 0) {
              setvalue((value: number) => +value - 1);
            }
          }}
          alt=""
          className={cls(styles["blue-arrow"], styles["left-blue-arrow"])}
        />
        <input type="number" value={value} onChange={handleInputChange} />
        <img
          onClick={() => {
            if (+value < count) {
              setvalue((value: number) => +value + 1);
            }
          }}
          src={blueArrow}
          alt=""
          className={cls(styles["blue-arrow"], styles["right-blue-arrow"])}
        />
      </div>
      <div className={styles["ticket-tip"]}>{t("请选择抽奖次数")}</div>
      <img src={splitLine} alt="" className="split-line" />
      <div className={styles["ticket-area"]}>
        <div className={styles["ticket-left"]}>
          <img src={veLP} alt="" />
          <span>{ticketNumber}</span>
        </div>
        <div className={styles["ticket-right"]} onClick={() => setvalue(count)}>
          <img src={maxText} alt="" className={styles["max-ticket"]} />
        </div>
      </div>
      <img src={splitLine} alt="" className="split-line" />
      <div className={styles["draw-button"]} onClick={handlSubmit}>
        <div>
          {i18n.language !== "kr" && t("Draw-with")}
          <img src={veLP} alt="" />
          <span>{value * Times}</span>
          {t("Draw")}
        </div>
      </div>
    </Modal>
  );
};

export default Index;
