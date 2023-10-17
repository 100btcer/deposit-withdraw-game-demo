export const TokenName = "gashapon_token";

const setLocal = (keyName: string) => (keyValue: string) => {
  window.localStorage.setItem(keyName, keyValue);
};

const getLocal = (keyName: string) => {
  return window.localStorage.getItem(keyName);
};

const removeLocal = (keyName: string) => {
  window.localStorage.removeItem(keyName);
};

export const setLocalToken = setLocal(TokenName);

export const getLocalToken = () => getLocal(TokenName);

export const removeLocalToken = () => removeLocal(TokenName);
