import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';

export class State {
  public id = 0; // 优惠券ID
  public couponName = ''; // 优惠券名称
  public nameLanguage = null; // 优惠券名称-多语言
  public type = 'reward'; // 优惠券类型 reward-满减 discount-折扣 random-随机
  public money = 0; // 发放面额 当type为reward时需要添加
  public discount = 1; // 1 =< 折扣 <= 9.9 当type为discount时需要添加
  public discountLimit = 0; // 最多折扣金额 当type为discount时可选择性添加
  public atLeast = 0; // 满多少元使用 0代表无限制
  public isShow = 1; // 是否显示，默认允许直接领取
  public count = 0; // 发放数量
  public maxFetch = 0; // 每人最大领取个数
  public validityType = 1; // 过期类型1-古固定时间范围过期 2-领取之日固定日期后过期 3长期有效
  public endUseTime = ''; // 使用结束日期 过期类型0时必填
  public fixedTerm = 0; // 当validity_type为2时需要添加 领取之日起或者次日N天内有效
  public scene = 1;// 场景 1-住宿 2-餐饮
  public propertySelectType = 1; // 1-全部物业 2-部分物业
  public propertyIds = ''; // 物业ID
  public propertyIdsArr = [];
  public restaurantSelectType = 1; // 1-全部餐厅 2-部分餐厅
  public restaurantIds = ''; // 餐厅IDs
  public restaurantIdsArr = [];
  public serviceIds = ''; // 服务IDs
  public serviceIdsArr = [];
  public carServiceTypes = ''; // 汽车服务类型
  public carServiceTypesArr = [];
  public leadCount = 0; // 已领取数量
  public usedCount = 0; // 已使用数量
  public startUseTime = 0; // 使用开始日期 过期类型0时必填
  public sort = 0; // 排序
  public discountAppStayOrderMoney = null; // 住宿订单的优惠总金额
  public appStayOrderMoney = null; // 住宿订单用券总成交额
  public status = 1; // 状态（1进行中2已结束-1已关闭）
  public createAt = ''; // 创建时间
  public updateAt = ''; // 修改时间
  public deletedAt = ''; // 删除时间
  public desc = ''; // 优惠券使用说明
  public descLanguage = null; // 优惠券使用说明

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


// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'type',
    component: 'NSelect',
    label: '优惠券类型',
    componentProps: {
      placeholder: '请输入优惠券类型',
      options: [
        {
          labelField: '满减',
          valueField: 'reward',
        },
        {
          labelField: '折扣',
          valueField: 'discount',
        }
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'couponName',
    component: 'NInput',
    label: '优惠券名称',
    componentProps: {
      placeholder: '请输入优惠券名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'validityType',
    component: 'NSelect',
    label: '有效期限',
    componentProps: {
      placeholder: '请选择',
      options: [
        {
          labelField: '固定时间',
          valueField: '1',
        },
        {
          labelField: '相对时间',
          valueField: '2',
        },
        {
          labelField: '长期有效',
          valueField: '3',
        }
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'scene',
    component: 'NSelect',
    label: '适用场景',
    componentProps: {
      placeholder: '请选择',
      options: [
        {
          labelField: '民宿',
          valueField: 1,
        },
        {
          labelField: '订餐',
          valueField: 2,
        },
        {
          labelField: '按摩',
          valueField: 3,
        },
        {
          labelField: '接送机/包车',
          valueField: 4,
        },
        {
          labelField: '储物柜',
          valueField: 5,
        },
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    componentProps: {
      placeholder: '请选择',
      options: [
        {
          labelField: '进行中',
          valueField: '1',
        },
        {
          labelField: '已结束',
          valueField: '2',
        },
        {
          labelField: '已关闭',
          valueField: '-1',
        }
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 90,
  },
  {
    title: '优惠券名称',
    key: 'couponName',
    align: 'left',
    width: 200,
  },
  {
    title: '券类型',
    key: 'type',
    align: 'left',
    width: 90,
    render(record) {
      if(record.type == 'reward'){
        return h(
          'div',
          {
            style: {
              color: '#17A158',
              width: '38px',
              height: '22px',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
              border: '1px solid #17A158'
            },
            bordered: false,
          },
          {
            default: () => '满减',
          }
        );
      }
      if(record.type == 'discount'){
        return h(
          'div',
          {
            style: {
              color: '#F48720',
              width: '38px',
              height: '22px',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
              border: '1px solid #F48720'
            },
            bordered: false,
          },
          {
            default: () => '折扣',
          }
        );
      }
      return '--'
    }
  },
  {
    title: '优惠内容',
    key: 'couponName',
    align: 'left',
    width: 180,
    ellipsis:false,
    render(record) {
      if(record.type=='reward'){
        return "满"+record.atLeast + "JPY减" + record.money + "JPY";
      }else if(record.type == 'discount'){
        return "满"+record.atLeast + "JPY打" + record.discount + "折；\n" + "最多可抵" + record.discountLimit + "JPY";
      }
    }
  },
  /*{
    title: '优惠券金额/折扣',
    key: 'couponName',
    align: 'left',
    width: 140,
    render(record) {
      if(record.type=='reward'){
        return record.money + "JPY";
      }else if(record.type == 'discount'){
        return record.discount + "折";
      }
    }
  },
  {
    title: '满多少元使用',
    key: 'atLeast',
    align: 'left',
    width: 120,
  },*/
  {
    title: '发放数量',
    key: 'count',
    align: 'left',
    width: 100,
    render(record) {
      if(record.count==0){
        return "无限制";
      }
      return record.count
    }
  },
  {
    title: '已领取数量',
    key: 'leadCount',
    align: 'left',
    width: 100,
  },
  {
    title: '已使用数量',
    key: 'usedCount',
    align: 'left',
    width: 100,
  },
  /*{
    title: '排序',
    key: 'sort',
    align: 'left',
    width: -1,
  },*/
  {
    title: '领取限制',
    key: 'maxFetch',
    align: 'left',
    width: 110,
    render(record) {
      if(record.maxFetch>0){
        return record.maxFetch + "张/人";
      }else{
        return '无领取限制';
      }
    }
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 100,
    render(record){
      if(record.status == 1){
        return h(
          'div',
          {
            style: {
              color: '#26A763',
              width: '52px',
              height: '22px',
              background: '#E3F4EB',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '进行中',
          }
        );
      }
      if(record.status == 2){
        return h(
          'div',
          {
            style: {
              color: '#EFA020',
              width: '52px',
              height: '22px',
              background: '#FFECCE',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '已结束',
          }
        );
      }
      if(record.status == -1){
        return h(
          'div',
          {
            style: {
              color: '#F56C6C',
              width: '52px',
              height: '22px',
              background: '#FEF0F0',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '已关闭',
          }
        );
      }
      return '--'
    }
  },
  {
    title: '适用场景',
    key: 'scene',
    align: 'left',
    width: 100,
    render(record) {
      if(record.scene == 1){
        return h(
          'div',
          {
            style: {
              color: '#26A763',
              padding: '0 5px',
              height: '22px',
              background: '#E3F4EB',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '民宿',
          }
        );
      }else if(record.scene == 2){
        return h(
          'div',
          {
            style: {
              color: '#3F9EFF',
              padding: '0 5px',
              height: '22px',
              background: '#ECF5FF',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '订餐',
          }
        );
      }else if(record.scene == 3){
        return h(
          'div',
          {
            style: {
              color: '#EFA020',
              padding: '0 5px',
              height: '22px',
              background: '#FFECCE',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '按摩',
          }
        );
      }else if(record.scene == 4){
        return h(
          'div',
          {
            style: {
              color: '#F56C6C',
              padding: '0 5px',
              height: '22px',
              background: '#FEF0F0',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '接送机/包车',
          }
        );
      }else if(record.scene == 5){
        return h(
          'div',
          {
            style: {
              color: '#F56C6C',
              padding: '0 5px',
              height: '22px',
              background: '#FEF0F0',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '储物柜',
          }
        );
      }
      return '--'
    }
  },
  {
    title: '创建时间',
    key: 'createAt',
    align: 'left',
    width: 155,
  },
  {
    title: '修改时间',
    key: 'updateAt',
    align: 'left',
    width: 155,
  },
];


