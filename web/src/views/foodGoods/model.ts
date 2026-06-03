import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import {NEllipsis, NTag} from 'naive-ui';
import {isNullObject} from "@/utils/is";

export class State {
  public id = 0; // id
  public restaurantId = 0; // 餐厅ID
  public toretaCourseId = ''; // Toreta的课程id
  public goodsName = ''; // 套餐名称
  public nameLanguage = null;
  public introduction = ''; // 促销语
  public labelIds = ''; // 标签（多选）
  public labelIdsArr = [];
  public images = ''; // 图集
  public imagesArr = [];
  public price = 0; // 套餐售价
  public marketPrice = 0; // 套餐原价
  public maxOrderNum = 0; // 单次最大可预定数
  public timeDuration = 0; // 用餐停留时间
  public goodsContent = ''; // 套餐详情
  public descriptionLanguage = null;
  public notice = ''; // 注意事项
  public goodsState = 1; // 状态（1.正常2下架）
  public isThCouponExclusive = 0; // 是否是礼品券兑换专属：0-否，1-是
  public thCouponId = null; // 绑定的礼品券ID，用于礼品券兑换专属商品
  public isNoPay = 0; // 是否是无需支付：0-否，1-是
  public maxTimeOrderOpen = 2; // 是否开启时间限制：2-否，1-是
  public orderTimeForm = '';// 时间限制
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
  public deletedAt = ''; // deleted_at

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则
export const rules = {
  price: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入套餐售价',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'goodsName',
    component: 'NInput',
    label: '套餐名称',
    componentProps: {
      placeholder: '请输入套餐名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'priceRange',
    component: 'NInput',
    label: '价格区间',
    componentProps: {
      pair: true,
      separator: '-',
      placeholder: ['最低价格', '最高价格'],
      min: 0,
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '套餐名称',
    key: 'goodsName',
    align: 'left',
    width: 150,
  },
  {
    title: '套餐详情',
    key: 'goodsContent',
    align: 'left',
    width: 400,
    ellipsis: false,
    render(row){
      return h(
        NEllipsis,
        {
          expandTrigger: 'click',
          lineClamp: 2,
          tooltip: false,
        },
        {
          default: () => row.goodsContent,
        }
      )
    }
  },
  {
    title: '套餐价格',
    key: 'price',
    align: 'left',
    width: 150,
    render(row){
      if(row.marketPrice){
        return row.price + 'JPY' + ' / ' + row.marketPrice + 'JPY'
      }else{
        return row.price + 'JPY'
      }
    }
  },
  {
    title: '单次最大可预约数',
    key: 'maxOrderNum',
    align: 'left',
    width: 150,
  },
  {
    title: '用餐停留时间',
    key: 'timeDuration',
    align: 'left',
    width: 140,
    render(row){
      if(row.timeDuration > 0){
        return row.timeDuration + '分钟'
      }else{
        return '无限制'
      }
    }
  },
  {
    title: '套餐状态',
    key: 'goodsState',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.goodsState)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.goodsState == 1 ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => row.goodsState == 1 ? '销售中' : '仓库中',
        }
      );
    },
  },
];


