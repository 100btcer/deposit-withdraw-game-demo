import { Modal } from "antd";
import { useState } from "react";
import type { FC } from "react";
import cls from "classnames";
import { useTranslation } from "react-i18next";
import styles from "./index.module.scss";
import closeIcon from "../../images/close-icon.png";
import GreenIcon from "../../images/tip-green.png";
import RedIcon from "../../images/tip-error.png";
import { scanUrl } from "../../conifg/scan";

interface ModalProsps {
  open: boolean;
  handleClose: () => void;
  tipType: "error" | "success";
  tipText: string;
  scanAddress?: string;
}

const Index: FC<ModalProsps> = ({
  open,
  handleClose,
  tipType,
  tipText,
  scanAddress,
}) => {
  const { t } = useTranslation();
  const imgSrc = tipType === "error" ? RedIcon : GreenIcon;
  return (
    <Modal
      wrapClassName={styles["body-wrrape"]}
      open={open}
      closable={false}
      footer={null}
      width={368}
      centered
    >
      <img
        src={closeIcon}
        alt=""
        className={styles["close-icon"]}
        onClick={handleClose}
      />
      <img src={imgSrc} alt="" className={styles["tip-icon"]} />
      <div className={styles["tip-text"]}>{tipText}</div>
      {scanAddress && (
        <a
          href={`${scanUrl}${scanAddress}`}
          target="_blank"
          rel="noopener noreferrer"
        >
          <div className={styles["tip-button"]}>{t("view-on-scane")}</div>
        </a>
      )}
    </Modal>
  );
};

export default Index;
