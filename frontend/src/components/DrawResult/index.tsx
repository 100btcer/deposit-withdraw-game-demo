import { useEffect, useState, useMemo } from "react";
import cls from "classnames";
import styles from "./index.module.scss";
import light1 from "../../images/light1.png";
import light2 from "../../images/light2.png";
import light3 from "../../images/light3.png";
import { useTranslation } from "react-i18next";
import mix from "../../images/Mix.png";

const Item = ({
  amount = 5000,
  img = mix,
  name = "MIX",
  type = 1,
  token_id = 1,
}) => {
  const { t } = useTranslation();
  const showName = useMemo(() => {
    if (+token_id === 7) {
      return t("egg-normal");
    }
    if (+token_id === 8) {
      return t("egg-privte");
    }
    return name;
  }, [name, token_id, t]);
  return (
    <div className={styles["content-item"]}>
      <div className={styles["content-top"]}>
        <img src={img} alt="" />
        <div className={styles["content-des"]}>{showName}</div>
      </div>
      <div className={styles["content-bottom"]}>
        {type === 1 ? amount : `x${amount}`}
      </div>
    </div>
  );
};

const Index = ({ handleCloseResult, list }: any) => {
  const [black, setBlack] = useState(false);
  const { t } = useTranslation();
  useEffect(() => {
    setBlack(true);
  }, []);
  return (
    <div
      className={cls(styles["container"], black && styles["black-container"])}
    >
      <div className={cls(styles["content"], black && styles["content-show"])}>
        <div className={styles["title"]}>{t("congratulation")}</div>
        <div className={styles["content-list"]}>
          {list &&
            list.map((result: any, index: number) => (
              <Item
                key={index}
                token_id={result.token_id}
                img={result.icon}
                amount={result.amount}
                name={result.name}
                type={result.type}
              />
            ))}
        </div>
        <div className={styles["view-button"]} onClick={handleCloseResult}>
          {t("view")}
        </div>
      </div>

      <img
        src={light1}
        alt=""
        className={cls(styles["light"], black && styles["light1"])}
      />
      <img
        src={light2}
        alt=""
        className={cls(styles["light"], black && styles["light2"])}
      />
      <img
        src={light3}
        alt=""
        className={cls(styles["light"], black && styles["light3"])}
      />
    </div>
  );
};

export default Index;
