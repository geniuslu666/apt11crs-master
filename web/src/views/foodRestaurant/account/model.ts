import { h, ref } from 'vue';
import {NTag} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public restaurantId = 0; // 餐厅ID
  public type = 'BANKCARD'; // 类型
  public bankName = ''; // 开户行
  public bankUser = ''; // 开户人姓名
  public bankCard = ''; // 银行账号
  public wechatAccount = ''; // 微信名
  public alipayName = ''; // 支付宝真实姓名
  public alipayAccount = ''; // 支付宝账户
  public status = 1; // 1、启用 2、禁用
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



// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '类型',
    key: 'type',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.type)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.account_type, row.type),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.account_type, row.type),
        }
      );
    },
  },
  {
    title: '账户信息',
    key: 'bankName',
    align: 'left',
    width: -1,
    render(row){
      if(row.type == 'BANKCARD'){
        return h(
          'div',
          {},
          [
            h(
              'div',
              {},
              {
                default: () => '开户行：'+row.bankName,
              }
            ),
            h(
              'div',
              {},
              {
                default: () => '开户人姓名：'+row.bankUser,
              }
            ),
            h(
              'div',
              {},
              {
                default: () => '银行账户：'+row.bankCard,
              }
            )
          ]
        )
      }else if(row.type == 'WECHAT'){
        return h(
          'div',
          {},
          [
            h(
              'div',
              {},
              {
                default: () => '微信名：'+row.wechatAccount,
              }
            ),
          ]
        )
      }else if(row.type == 'ALIPAY'){
        return h(
          'div',
          {},
          [
            h(
              'div',
              {},
              {
                default: () => '支付宝真实姓名：'+row.alipayName,
              }
            ),
            h(
              'div',
              {},
              {
                default: () => '支付宝账户：'+row.alipayAccount,
              }
            ),
          ]
        )
      }else if(row.type == 'CASH'){
        return '--'
      }else if(row.type == 'OTHER'){
        return '--'
      }
    }
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
        }
      );
    },
  },
  {
    title: '创建时间',
    key: 'createAt',
    align: 'left',
    width: -1,
  },
  {
    title: '更新时间',
    key: 'updateAt',
    align: 'left',
    width: -1,
  },
];

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
  account_type: [] as Option[]
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable','account_type'],
  }).then((res) => {
    options.value = res;
  });
}


