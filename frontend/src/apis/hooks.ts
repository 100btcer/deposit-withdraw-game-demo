import { useMutation, useQuery } from "react-query";
import * as apis from "./api";

/**
 *
 *  @description 投放部审核物料
 */
export const useMutationLogin = () =>
  useMutation({
    mutationFn: apis.loginRequest,
  });

/**
 * @description 用户信息
 *
 */
export const useRequestUserInfo = () => {
  const query = useQuery(["User-info"], () => apis.getUserInfo());
  return query;
};

/**
 * @description 获取奖品列表
 *
 */
export const useRequestRewardList = (params: { type?: 1 | 2 }) => {
  const query = useQuery(
    [`/lottery/result/list${params.type}`],
    () => apis.getRewardList({ page: 1, page_size: 1000, type: params.type }),
    {
      staleTime: 0,
    }
  );
  return query;
};

/**
 * @description 奖品列表
 *
 */
export const useRequestRewardResult = (hash: string) => {
  const query = useQuery(
    [`asdasds${hash}`],
    () => apis.getLotteryResult({ hash }),
    { enabled: !!hash, refetchInterval: 1000 }
  );
  return query;
};

/**
 * @description 获取质押记录
 *
 */
export const useRequestDepositeList = () => {
  const query = useQuery([`/api/deposit/record/list`], () =>
    apis.getDepositeRecordList({ page: 1, page_size: 1000 })
  );
  return query;
};

export const useMutationGetLottery = () =>
  useMutation({
    mutationFn: apis.getLottery,
  });

export const useRequestGetValideNumber = () => {
  const query = useQuery([`getvalid_amount`], () => apis.getvalid_amount());
  return query;
};
