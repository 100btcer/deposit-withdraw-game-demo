import "../App.css";
import Content from "../components/Content";
import "@rainbow-me/rainbowkit/styles.css";
import { getDefaultWallets, RainbowKitProvider } from "@rainbow-me/rainbowkit";
import { QueryClient, QueryClientProvider } from "react-query";
import { publicProvider } from "wagmi/providers/public";

import { rangers } from "../conifg/chain";
import { configureChains, createConfig, WagmiConfig } from "wagmi";
const chain = [rangers];
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      //  1 秒内相同的接口不会重新请求
      staleTime: 1 * 1000,
      retry: 0,
    },
  },
});

const { chains, publicClient } = configureChains(chain, [publicProvider()]);
const { connectors } = getDefaultWallets({
  appName: "My RainbowKit App",
  projectId: "3e69fa9619363d40466be3c497843b4f",
  chains,
});
const wagmiConfig = createConfig({
  autoConnect: true,
  connectors,
  publicClient,
});

function App({ handlePlay, musicOpen }: any) {
  return (
    <QueryClientProvider client={queryClient}>
      <WagmiConfig config={wagmiConfig}>
        <RainbowKitProvider chains={chains}>
          <Content handlePlay={handlePlay} musicOpen={musicOpen} />
        </RainbowKitProvider>
      </WagmiConfig>
    </QueryClientProvider>
  );
}

export default App;
