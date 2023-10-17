import request from "../utils/request";
import { getLocalToken } from "../utils/localStorage";

interface CommonParams {
  page: number;
  page_size: number;
}

interface DepositeRecordItem {
  id: number;
  user_id: number;
  lock_day: number;
  create_time: string;
  deposit_amount: number;
  hash: string;
}

interface RewardListItem {
  name: string;
  type: 1 | 2;
  amount: number;
  link: string;
  token_id: number;
  icon: string;
}
type CommonListReturn<T> = {
  total: number;
  list: Array<T>;
};
type DepositeRecordReturn = CommonListReturn<DepositeRecordItem>;

type RewardListReturn = CommonListReturn<RewardListItem>;

export const loginRequest = (params: {
  address: string;
  message: string;
  sign: string;
}) =>
  request.post("/api/user/wallet/login", params) as Promise<{
    token: string;
    user_id: number;
  }>;

/**
 *
 * @descript 查询质押记录
 *
 */
export const getDepositeRecordList = (params: CommonParams) => {
  if (!getLocalToken()) {
    return;
  }
  return request.get(
    "/api/deposit/record/list",
    params
  ) as Promise<DepositeRecordReturn>;
};

/**
 *
 * @descript 获取奖品列表
 *
 */
export const getRewardList = (params: CommonParams & { type?: 1 | 2 }) => {
  if (!getLocalToken()) {
    return;
  }
  return request.get(
    "/api/lottery/result/list",
    params
  ) as Promise<RewardListReturn>;
};

/**
 *
 * @descript 获取用户信息
 *
 */
export const getUserInfo = () => {
  if (!getLocalToken()) {
    return;
  }
  return request.get("/api/user/info") as Promise<{
    user_id: number;
    address: string;
    deposit_amount: number;
  }>;
};

/**
 *
 * @descript 获取抽奖结果
 *1
 */
export const getLotteryResult = (params: { hash: string }) => {
  if (!getLocalToken()) {
    return;
  }
  return request.get("/api/lottery/result/get", params) as Promise<
    Array<RewardListItem>
  >;
};

/**
 *
 * @descript 领奖
 *1
 */
export const getLottery = (params: { id: string }) => {
  return request.get("/api/lottery/prize/award/giveout", params) as Promise<
    Array<unknown>
  >;
};

export const getvalid_amount = () => {
  if (!getLocalToken()) {
    return;
  }
  return request.get("/api/withdraw/valid_amount/get") as Promise<any>;
};
