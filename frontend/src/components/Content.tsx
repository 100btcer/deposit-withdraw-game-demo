import { useEffect, useMemo, useState, useRef, lazy } from "react";
import styles from "./Content.module.scss";
import { useTranslation } from "react-i18next";
import { useWeb3Modal } from "@web3modal/react";
import {
  useConnectModal,
  useAccountModal,
  useChainModal,
} from "@rainbow-me/rainbowkit";
import {
  useAccount,
  useSignMessage,
  useContractRead,
  useNetwork,
  useSwitchNetwork,
} from "wagmi";
import erthers from "ethers";
import type { MenuProps } from "antd";
import { Howl } from "howler";
import { formatUnits } from "ethers";
import cls from "classnames";
import {
  modal_sound,
  other_buton_sound,
  draw_button_sound,
} from "../conifg/sound";
import { erc20ABI } from "wagmi";
import { Dropdown } from "antd";

import {
  useMutationLogin,
  useRequestRewardResult,
  useRequestGetValideNumber,
} from "../apis/hooks";
import {
  setLocalToken,
  getLocalToken,
  removeLocalToken,
} from "../utils/localStorage";
import walletAddress from "../images/wallet-address.png";
import cnIcon from "../images/cn-icon.png";
import krIcon from "../images/kr-icon.png";
import enIcon from "../images/en-icon.png";
import helpImage from "../images/help.png";
import RangersCoin from "../images/ranger-small.png";
import music from "../images/music.png";
import musicClose from "../images/music-close.png";
import bgImage from "../images/bg-image.png";
import { lpAddress, ticketAddress } from "../conifg/scan";
import ticketAbi from "../../src/conifg/ticketabi.json";
import DrawModal from "./DrawModal";
import Mymodal from "./Modal";
import DrawCn from "../../src/images/langImages/draw-cn.png";
import DrawEn from "../../src/images/langImages/draw-en.png";
import DrawKr from "../../src/images/langImages/draw-kr.png";

const drawTextMap: Record<string, string> = {
  kr: DrawKr,
  en: DrawEn,
  cn: DrawCn,
};

const BagModal = lazy(
  /* webpackChunkName: "BagModal"*/
  /* webpackPrefetch: true */
  /* webpackPreload: true */
  () => import("./BagModal")
);

const TipModal = lazy(
  /* webpackChunkName: "TipModal"*/
  /* webpackPrefetch: true */
  /* webpackPreload: true */
  () => import("./TipModal")
);

const RrawResult = lazy(
  /* webpackChunkName: "RrawResult"*/
  /* webpackPrefetch: true */
  /* webpackPreload: true */
  () => import("./DrawResult")
);
const signMesage = "GASHAPON";

const allLangItems = [
  {
    key: "cn",
    label: "CN",
  },
  {
    key: "kr",
    label: "KR",
  },
  {
    key: "en",
    label: "EN",
  },
];

const VideoPreloader = () => {
  return (
    <video
      src="https://starlands3.s3.ap-southeast-1.amazonaws.com/starland/1689047737212-drawwating+.mp4"
      preload="auto"
      autoPlay
      loop
      muted
      playsInline
      style={{ display: "none" }}
    />
  );
};

export const handleFormatList = (list: any[]) => {
  const a: any = {};
  list.forEach((item) => {
    if (a[item.token_id]) {
      a[item.token_id].amount = a[item.token_id].amount + item.amount;
    } else {
      a[item.token_id] = item;
    }
  });
  const tempList = Object.keys(a).map((v) => a[v]);
  return tempList;
};

const langIconMap = {
  en: enIcon,
  kr: krIcon,
  cn: cnIcon,
};
export const network = process.env.REACT_APP_FIRST_ENV === "prod" ? 2025 : 9527;
export const useChangeNetwork = () => {
  const { chain } = useNetwork();
  const { switchNetworkAsync } = useSwitchNetwork();
  const networkError = useMemo(() => {
    if (chain && chain.id !== network) {
      return true;
    }
    return false;
  }, [chain]);

  const handleChangeNetwork = () => {
    if (networkError) {
      return switchNetworkAsync?.(network);
    }
  };

  return { networkError, handleChangeNetwork };
};

const Index = ({
  handlePlay,
  musicOpen,
}: {
  handlePlay: () => void;
  musicOpen: boolean;
}) => {
  const { t, i18n } = useTranslation("translation");
  const isfirstRender = useRef<boolean>(true);
  const [videoLoaded, setVideoLoaded] = useState(false);
  const [showWaitvideo, setShowWaitvideo] = useState(false);
  const [preloadVideo, setPreloadVideo] = useState(false);
  const [rewardHash, setRewardHash] = useState("");
  const language = i18n.language as unknown as keyof typeof langIconMap;
  const { openConnectModal } = useConnectModal();
  const { openAccountModal } = useAccountModal();
  const modalHowl = useRef<any>(null);
  const otherHowl = useRef<any>(null);
  const drawButtonHowl = useRef<any>(null);
  const [showDeposite, setSHowDeposite] = useState<boolean>(false);
  const [showBag, setSHowBag] = useState<boolean>(false);
  const [showDraw, setShowDraw] = useState<boolean>(false);
  const [showResult, setShowResult] = useState<boolean>(false);
  const [resultList, setResultList] = useState<any>([]);
  const { address, isConnected, connector } = useAccount();
  const { networkError } = useChangeNetwork();
  const { mutateAsync: loginAysnc } = useMutationLogin();
  const { data: dataAmount, refetch } = useRequestGetValideNumber();
  const { data: rewardResult } = useRequestRewardResult(rewardHash);
  useEffect(() => {
    if (rewardResult && rewardResult.length > 0) {
      setRewardHash("");
      handleOpenResult();
      const res = handleFormatList(rewardResult);
      setResultList(res);
    }
  }, [rewardResult]);
  useEffect(() => {
    if (address && !isfirstRender) {
      removeLocalToken();
    }
  }, [address]);
  const { signMessageAsync, error } = useSignMessage({
    message: signMesage,
  });
  const { data: ticketdata, refetch: TiketNumberRefetch } = useContractRead({
    address: ticketAddress,
    abi: ticketAbi,
    functionName: "balanceOf",
    args: [address || "0x"],
  });
  const { data: tokendata, refetch: lpNumberRefecth } = useContractRead({
    address: lpAddress,
    abi: erc20ABI,
    functionName: "balanceOf",
    args: [address || "0x"],
  });
  const lpNumber = useMemo(() => {
    if (tokendata) {
      return parseFloat((+formatUnits(tokendata?.toString(), 18)).toFixed(4));
    }
    return 0;
  }, [tokendata]);
  const ticketNumber = useMemo(() => {
    if (ticketdata) {
      console.log(
        Math.floor(+formatUnits(ticketdata?.toString(), 18)),
        "ticketdata"
      );
      return Math.floor(+formatUnits(ticketdata?.toString(), 18));
    }
    return 0;
  }, [ticketdata]);
  const items: MenuProps["items"] = [
    {
      key: "1",
      label: (
        <a
          target="_blank"
          rel="noopener noreferrer"
          href="https://ponytaswap.gitbook.io/ponytaswap/"
        >
          {t("Tutorial")}
        </a>
      ),
    },
    {
      key: "2",
      label: (
        <a
          target="_blank"
          rel="noopener noreferrer"
          href="https://bridge.rangersprotocol.com"
        >
          {t("Brige")}
        </a>
      ),
    },
    {
      key: "3",
      label: (
        <a
          target="_blank"
          rel="noopener noreferrer"
          href="https://www.ponytaswap.finance/liquidity"
        >
          {t("Get LP")}
        </a>
      ),
    },
  ];

  useEffect(() => {
    modalHowl.current = new Howl({
      src: modal_sound,
      loop: false,
      html5: true,
    });
  }, []);
  useEffect(() => {
    otherHowl.current = new Howl({
      src: other_buton_sound,
      loop: false,
      html5: true,
    });
  }, []);
  useEffect(() => {
    drawButtonHowl.current = new Howl({
      src: draw_button_sound,
      loop: false,
      html5: true,
    });
  }, []);
  useEffect(() => {
    if (
      isConnected &&
      !isfirstRender.current &&
      connector &&
      !getLocalToken()
    ) {
      login();
    }
    isfirstRender.current = false;
  }, [isConnected]);
  useEffect(() => {
    if (!getLocalToken() && isConnected && connector) {
      login();
    }
  }, [connector]);
  useEffect(() => {
    if (!address) {
      removeLocalToken();
    }
  }, [address]);
  const login = async () => {
    const res = await signMessageAsync();
    const loginres = await loginAysnc({
      sign: res,
      address: address!,
      message: signMesage,
    });
    setLocalToken(loginres.token);
  };
  const showAddress = useMemo(() => {
    if (address) {
      const first = address.slice(0, 6);
      const last = address.slice(address.length - 5, address.length);
      return `${first}...${last}`;
    }
    return "";
  }, [address]);
  const handlePlayOtherButtonSound = () => {
    otherHowl.current.play();
  };
  const langItems = useMemo(() => {
    return allLangItems.filter((v) => v.key !== language);
  }, [language]);

  const handleChangeLanguae: MenuProps["onClick"] = ({ key }) => {
    i18n.changeLanguage(key);
  };
  const handleChangeMusic = () => {
    handlePlay();
  };
  const handleConnect = async () => {
    openConnectModal?.();
  };
  const handlePlayDrawButton = () => {
    drawButtonHowl.current.play();
  };

  const handleOpenResult = () => {
    setShowResult(true);
  };

  const handleCloseResult = () => {
    setShowResult(false);
    setSHowBag(true);
    setShowWaitvideo(false);
    TiketNumberRefetch();
  };
  const lang = i18n.language;
  return (
    <>
      {showResult && (
        <RrawResult
          handleCloseResult={handleCloseResult}
          list={resultList}
        ></RrawResult>
      )}

      {showWaitvideo ? (
        <video
          src="https://starlands3.s3.ap-southeast-1.amazonaws.com/starland/1689047737212-drawwating+.mp4"
          loop
          muted
          autoPlay
          playsInline
          className={cls(styles["bg-video"])}
        ></video>
      ) : (
        <video
          autoPlay
          loop
          muted
          playsInline
          onCanPlayThrough={() => {
            setPreloadVideo(true);
            setVideoLoaded(true);
          }}
          poster={bgImage}
          src="https://starlands3.s3.ap-southeast-1.amazonaws.com/starland/1689140262837-bg.mp4"
          className={styles["bg-video"]}
        ></video>
      )}
      <VideoPreloader></VideoPreloader>
      <div>
        <div className="header">
          <div>
            {address ? (
              <div
                className={styles["wallet-address"]}
                onClick={() => {
                  openAccountModal?.();
                }}
              >
                <img src={walletAddress} alt="" />
                {networkError ? "Network Error" : showAddress}
              </div>
            ) : (
              <div className={styles["wallet-connect-button"]}>
                <div
                  className={styles["wallet-connect-inner"]}
                  onClick={handleConnect}
                >
                  <span
                    style={
                      lang === "en"
                        ? {
                            transform: "scale(0.78) translateX(-7px)",
                            display: "inline-block",
                          }
                        : {}
                    }
                  >
                    {t("wallet-connect")}
                  </span>
                </div>
              </div>
            )}
          </div>
          {address && (
            <div className={styles["header-center"]}>
              <img src={RangersCoin} alt="" />
              <span>{lpNumber}</span>
            </div>
          )}
        </div>
        <div className={styles["header-right"]}>
          <Dropdown
            menu={{ items: langItems, onClick: handleChangeLanguae }}
            placement="bottom"
            trigger={["click"]}
            getPopupContainer={(el) => el.parentElement || el}
          >
            <img
              src={langIconMap[language] || langIconMap["cn"]}
              alt=""
              className={styles["header-icon"]}
            />
          </Dropdown>
          <img
            src={musicOpen ? music : musicClose}
            alt=""
            onClick={handleChangeMusic}
            className={styles["header-icon"]}
          />
          <Dropdown
            menu={{ items }}
            placement="bottomRight"
            trigger={["click"]}
            getPopupContainer={(el) => el.parentElement || el}
          >
            <img src={helpImage} alt="" className={styles["header-icon"]} />
          </Dropdown>
        </div>
        {address && (
          <div
            className={styles["myReward-button"]}
            onClick={() => {
              modalHowl.current.play();
              setSHowBag(true);
            }}
          >
            <div className={styles["inner-text"]}> {t("my-reward")}</div>
          </div>
        )}
        <div
          className={styles["deposite-button"]}
          onClick={() => {
            modalHowl.current.play();
            if (address) {
              setSHowDeposite(true);
            } else {
              handleConnect();
            }
          }}
        >
          <div className={styles["inner-text"]}>
            <div>{t("deposite")}</div>
            <div>{t("getTicket")}</div>
          </div>
        </div>
        {!showWaitvideo && (
          <div
            className={styles["draw-button"]}
            onClick={() => {
              drawButtonHowl.current.play();
              if (address) {
                setShowDraw(true);
              } else {
                handleConnect();
              }
            }}
          >
            <img
              src={drawTextMap[lang] || DrawEn}
              alt=""
              className={cls(styles["draw-text"])}
            />
          </div>
        )}
        {!showWaitvideo && (
          <div className={styles["count-ticket"]}>
            <div className={styles["count-ticket-text"]}>
              {ticketNumber
                .toString()
                .padStart(7, "0")
                .split("")
                .map((a, index) => (
                  <span key={index}>{a}</span>
                ))}
            </div>
          </div>
        )}
      </div>

      <Mymodal
        open={showDeposite}
        lpNumber={lpNumber}
        lpNumberRefecth={lpNumberRefecth}
        TiketNumberRefetch={TiketNumberRefetch}
        handleClose={() => setSHowDeposite(false)}
        handlePlayOtherButtonSound={handlePlayOtherButtonSound}
      />
      {showBag && (
        <BagModal
          open={showBag}
          handleClose={() => setSHowBag(false)}
          handlePlayOtherButtonSound={handlePlayOtherButtonSound}
        />
      )}
      {showDraw && (
        <DrawModal
          open={showDraw}
          ticketNumber={ticketNumber}
          handleShowWaiteVideo={(hash) => {
            setRewardHash(hash);
            setShowWaitvideo(true);
            TiketNumberRefetch();
          }}
          handleClose={() => {
            setShowDraw(false);
            TiketNumberRefetch();
          }}
          handlePlayDrawButton={handlePlayDrawButton}
        />
      )}
    </>
  );
};

export default Index;
