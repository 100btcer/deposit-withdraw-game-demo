import { useEffect, useRef, useState, Suspense, lazy } from "react";
import { Howl } from "howler";
import "./App.css";
import { bg_sound } from "./conifg/sound";

const Page = lazy(
  /* webpackChunkName: "Page"*/
  /* webpackPrefetch: true */
  /* webpackPreload: true */
  /* webpackMode: "lazy" */
  () => import("./components/Page")
);

function App() {
  const howl = useRef<any>();
  const FirstPlay = useRef<boolean>(false);
  const [musicOpen, setMusicOpen] = useState<boolean>(false);
  useEffect(() => {
    howl.current = new Howl({
      src: bg_sound,
      loop: true,
      html5: true,
      onplay: () => {
        FirstPlay.current = true;
      },
    });
  }, []);
  const handleFirstPlay = () => {
    if (!FirstPlay.current) {
      setMusicOpen(true);
      howl.current.play();
    }
  };
  const handlePlay = () => {
    if (musicOpen) {
      howl.current.pause();
    } else {
      howl.current.play();
    }
    setMusicOpen((a) => !a);
  };
  return (
    <div
      className="App"
      onClick={handleFirstPlay}
      onTouchStart={handleFirstPlay}
    >
      <div className="content">
        <Suspense
          fallback={
            <video
              src="https://starlands3.s3.ap-southeast-1.amazonaws.com/starland/1689140770489-Loading-Animation.mp4"
              loop
              muted
              autoPlay
              playsInline
              className="loading-video"
            ></video>
          }
        >
          <Page handlePlay={handlePlay} musicOpen={musicOpen}></Page>
        </Suspense>
      </div>
    </div>
  );
}

export default App;
