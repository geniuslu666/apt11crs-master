import { ThMchInter } from './thMch';

export interface ThCouponInter {
  id: number;
  couponName: string;
  mchList: ThCouponMchInter[];
}

export interface ThCouponMchInter {
  couponId: number;
  mchId: number;
  name: string;
  mchInfo: ThMchInter;
}


export type ThCouponList = ThCouponInter[];