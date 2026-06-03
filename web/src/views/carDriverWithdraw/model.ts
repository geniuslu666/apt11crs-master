import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // id
  public type = ''; // 类型
  public driverId = 0; // 司机ID
  public driverDetail = {
    name: ''
  }; // 司机信息
  public withdrawSn = ''; // 提现单号
  public withdrawStatus = 'WAIT'; // 提现状态
  public withdrawAmount = null; // 提现金额
  public arrivalAmount = null; // 到账金额
  public serviceCharge = null; // 提现手续费比例
  public transfer = 0; // 1、未转账 2、已转账
  public applyRemark = ''; // 审核备注
  public applyAt = ''; // 审核时间
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public deletedAt = ''; // 删除时间

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
  type: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入类型',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'withdrawSn',
    component: 'NInput',
    label: '提现单号',
    componentProps: {
      placeholder: '请输入提现单号',
      onUpdateValue: (e: any) => {
        // console.log(e);
      },
    },
  },
  {
    field: 'driverName',
    component: 'NInput',
    label: '司机名称',
    componentProps: {
      placeholder: '请输入司机名称',
      onUpdateValue: (e: any) => {
        // console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    label: '创建时间',
    slot: 'createdAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: '5%',
  },
  {
    title: '提现单号',
    key: 'withdrawSn',
    align: 'left',
    width: -1,
  },
  {
    title: '司机',
    key: 'driverId',
    align: 'left',
    width: -1,
    render(record) {
      return record.driverDetail.name
    }
  },
  {
    title: '提现状态',
    key: 'withdrawStatus',
    align: 'left',
    width: -1,
    render(record) {
      if(record.withdrawStatus == 'WAIT'){
        return h(
          'span',
          {
            style: {
              color: 'green',
            },
            bordered: false,
          },
          {
            default: () => '待审核',
          }
        );
      }
      if(record.withdrawStatus == 'SUCCESS'){
        return h(
          'span',
          {
            style: {
              color: 'blue',
            },
            bordered: false,
          },
          {
            default: () => '已同意',
          }
        );
      }
      if(record.withdrawStatus == 'FAIL'){
        return h(
          'span',
          {
            style: {
              color: 'red',
            },
            bordered: false,
          },
          {
            default: () => '已拒绝',
          }
        );
      }
      return '--'
    }
  },
  {
    title: '提现金额',
    key: 'withdrawAmount',
    align: 'left',
    width: -1,
  },
  {
    title: '到账金额',
    key: 'arrivalAmount',
    align: 'left',
    width: -1,
  },
  {
    title: '提现手续费比例',
    key: 'serviceCharge',
    align: 'left',
    width: -1,
  },
  {
    title: '转账状态',
    key: 'transfer',
    align: 'left',
    width: -1,
    render(record) {
      if(record.transfer == 1){
        return h(
          'span',
          {
            style: {
              color: 'green',
            },
            bordered: false,
          },
          {
            default: () => '待转账',
          }
        );
      }
      if(record.transfer == 2){
        return h(
          'span',
          {
            style: {
              color: 'blue',
            },
            bordered: false,
          },
          {
            default: () => '已转账',
          }
        );
      }
      return '--'
    }
  },
  {
    title: '审核备注',
    key: 'applyRemark',
    align: 'left',
    width: -1,
  },
  {
    title: '审核时间',
    key: 'applyAt',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 165,
  },

];


