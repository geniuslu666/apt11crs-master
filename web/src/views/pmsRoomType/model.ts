import {ref} from 'vue';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {defRangeShortcuts} from '@/utils/dateUtil';

export class State {
  public uid = ''; // 第三方系统的ID
  public id = 0; // 主键ID
  public puid = ''; // 物业ID
  public cover = ''; // 封面
  public coverList = []; // 图集
  public name = ''; // 房型名称
  public roomRatePlanInfos = null; // 房型费率计划
  public nameLanguage = null;
  public basePrice = null; // 最低价格
  public checkinAt = '00:00'; // 入住时间
  public checkoutAt = '00:00'; // 退房时间
  public bookingStyle = ''; // 预订方式
  public roomStyle = ''; // 房间风格
  public occupancy = 0; // 占用
  public size = ''; // 房间大小
  public bedrooms = null; // 卧室
  public bathrooms = ''; // 浴室
  public cleaningFee = null; // 清理费
  public additionalGuestAmounts = null; // 额外客人金额
  public occupantsForBaseRate = null; // 无需增加额外客人金额人数
  public createAt = ''; // create_at
  public updateAt = ''; // update_at
  public bedType = null; // 床型

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
  cover: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入封面',
  },
  // name: {
  //   required: true,
  //   trigger: ['blur', 'input'],
  //   type: 'string',
  //   message: '请输入房型名称',
  // },
  // checkinAt: {
  //   required: true,
  //   trigger: ['blur', 'input'],
  //   type: 'string',
  //   message: '请输入入住时间',
  // },
  // checkoutAt: {
  //   required: true,
  //   trigger: ['blur', 'input'],
  //   type: 'string',
  //   message: '请输入退房时间',
  // },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  // {
  //   field: 'puid',
  //   component: 'NInput',
  //   label: '物业ID',
  //   componentProps: {
  //     placeholder: '请输入物业ID',
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  {
    field: 'name',
    component: 'NInput',
    label: '房型名称',
    componentProps: {
      placeholder: '请输入房型名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'checkinAt',
    component: 'NDatePicker',
    label: '入住时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);


