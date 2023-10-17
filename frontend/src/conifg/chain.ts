import { Chain } from "wagmi";

console.log(process.env.REACT_APP_FIRST_ENV, "--22222-");

export const rangers =
  process.env.REACT_APP_FIRST_ENV === "prod"
    ? {
        id: 2025,
        name: "rangers",
        network: "rangers",
        nativeCurrency: {
          decimals: 18,
          name: "rangers",
          symbol: "RPG",
        },
        rpcUrls: {
          public: { http: ["https://mainnet.rangersprotocol.com/api/jsonrpc"] },
          default: {
            http: ["https://mainnet.rangersprotocol.com/api/jsonrpc"],
          },
        },
        blockExplorers: {
          etherscan: {
            name: "SnowTrace",
            url: "https://scan.rangersprotocol.com",
          },
          default: {
            name: "SnowTrace",
            url: "https://scan.rangersprotocol.com",
          },
        },
      }
    : ({
        id: 9527,
        name: "robin",
        network: "robin",
        nativeCurrency: {
          decimals: 18,
          name: "robin",
          symbol: "RPG",
        },
        rpcUrls: {
          public: { http: ["https://robin.rangersprotocol.com/api/jsonrpc"] },
          default: { http: ["https://robin.rangersprotocol.com/api/jsonrpc"] },
        },
        blockExplorers: {
          etherscan: {
            name: "SnowTrace",
            url: "https://robin-rangersscan.rangersprotocol.com",
          },
          default: {
            name: "SnowTrace",
            url: "https://robin-rangersscan.rangersprotocol.com",
          },
        },
      } as const satisfies Chain);
