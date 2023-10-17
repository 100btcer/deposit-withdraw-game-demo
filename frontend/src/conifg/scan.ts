const isProd = process.env.REACT_APP_FIRST_ENV === "prod";

export const scanUrl = isProd
  ? "https://scan.rangersprotocol.com/tx/"
  : "https://robin-rangersscan.rangersprotocol.com/tx/";

export const lpAddress = isProd
  ? "0xd690fcc180913403a3ac7b19fd99a669f31b2f3f"
  : "0x2D2E7dC4b204f6b24eF0Cb719D473757651488ed";

export const ticketAddress = isProd
  ? "0xa04b3619f70c21cc7d2cc5c80729ae8d68f738be"
  : "0x2c3b9ce8a9b40ec39a7b0cc4d3ccf2374b2fb4b8";

export const mmAddress = isProd
  ? "0xAa8519721c49E87B88aAec826434368D4e7317D0"
  : "0x1deFB5b39B9dACBF75479347B34Dd8d7011f6901";
