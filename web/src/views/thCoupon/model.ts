import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import {defRangeShortcuts} from "@/utils/dateUtil";
import {isNullObject} from "@/utils/is";
import {NEllipsis, NImage, NTag} from "naive-ui";
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";
import {All} from "@/api/thCouponCategory";
import defaultImg from "@/assets/images/thcoupondf.png";

export class State {
  public id = 0; // 礼品券ID
  public logo = null;
  public couponName = ''; // 礼品券名称
  public nameLanguage = null; // 礼品券名称-多语言
  public couponSubName = ''; // 礼品券副标题
  public identityName = ''; // 券识别名称
  public subNameLanguage = null; // 礼品券副标题-多语言
  public couponNoPrefix = ''; // 编号前缀
  public status = 1; // 发放状态（1立即启用  2暂不启用）
  public fixedTerm = 0; // 领取之日起或者次日N天内有效
  public desc = ''; // 礼品券使用说明
  public descLanguage = null; // 礼品券使用说明
  public useStatus = 1; // 使用状态（1开始使用  2停止使用）
  public useMode = 'ARRIVE_VERIFY'; // 使用模式
  public categoryId = null; // 分类ID
  public mchList = []; // {mchId: 1, name: ''}
  public sort = 0; // 排序
  public createAt = ''; // 创建时间
  public updateAt = ''; // 修改时间
  public deletedAt = ''; // 删除时间
  public needReservation = 0; // 是否需要预约：0-不需要，1-需要
  public reservationRestaurantIds = ''; // 需要预约的餐厅IDs，多个用逗号分隔，如：1,2,3
  public reservationRestaurantIdsArr = [];

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
    field: 'couponName',
    component: 'NInput',
    label: '券名称',
    componentProps: {
      placeholder: '请输入礼品券名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'identityName',
    component: 'NInput',
    label: '券识别名称',
    componentProps: {
      placeholder: '请输入券识别名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'categoryId',
    component: 'NSelect',
    label: '券分类',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择礼品券分类',
      options: [],
      labelField: 'name',
      valueField: 'id',
      onUpdateValue: (e: any) => {
        console.log(typeof(e));
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '发行状态',
    componentProps: {
      placeholder: '请选择',
      options: [
        {
          labelField: '发行中',
          valueField: '1',
        },
        {
          labelField: '已停止发行',
          valueField: '2',
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
    field: 'useStatus',
    component: 'NSelect',
    label: '使用状态',
    componentProps: {
      placeholder: '请选择',
      options: [
        {
          labelField: '使用中',
          valueField: '1',
        },
        {
          labelField: '已停止使用',
          valueField: '2',
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
    field: 'createAt',
    label: '创建时间',
    slot: 'createAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: '礼品券名称',
    key: 'couponName',
    align: 'left',
    width: 200,
  },
  {
    title: '券识别名称',
    key: 'identityName',
    align: 'left',
    width: 180,
    // ellipsis: {
    //   tooltip: true
    // },
  },
  /*{
    title: '券LOGO',
    key: 'logo',
    align: 'left',
    width: 120,
    render(row){
      if(row.logo){
        return h(
          NImage,
          {
            width: 80,
            src: row.logo,
            onError: (e) => {
              e.target.src = defaultImg
            }
          },

        )
      }else{
        return h(
          NImage,
          {
            width: 80,
            src: defaultImg,
          },

        )
      }
    }
  },*/
  {
    title: '适用商户',
    key: 'mch',
    align: 'left',
    width: 200,
    ellipsis: false,
    render(row){
      if(row.mchList && row.mchList.length > 0){
        if(row.mchList.length > 1){
          var htmlArr = [];
          row.mchList.forEach((item)=>{
            var htmlItem = h(
              'div',
              null,
              {
                default: () => item.mchInfo?.name
              }
            )
            htmlArr.push(htmlItem)
          })
          return h(
            NEllipsis,
            {
              style:{
                width: '200px',
              },
              expandTrigger: 'click',
              lineClamp: 1,
              tooltip: false,
            },
            [
              htmlArr
            ]
          )
        }else{
          return row.mchList[0].mchInfo?.name
        }
      }else{
        return '--'
      }
    }
  },
  {
    title: '已发放数量',
    key: 'count',
    align: 'left',
    width: 100,
    render(record) {
      return record.count
    }
  },
  {
    title: '已使用数量',
    key: 'usedCount',
    align: 'left',
    width: 100,
  },
  {
    title: '发行状态',
    key: 'status',
    align: 'left',
    width: 110,
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
          type: getOptionTag(options.value.th_coupon_status, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.th_coupon_status, row.status),
        }
      );
    },
  },
  {
    title: '使用状态',
    key: 'useStatus',
    align: 'left',
    width: 110,
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
          type: getOptionTag(options.value.th_coupon_use_status, row.useStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.th_coupon_use_status, row.useStatus),
        }
      );
    },
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

export const cateList = ref([]);
// 字典数据选项
export const options = ref({
  th_coupon_status: [] as Option[],
  th_coupon_use_status: [] as Option[],
});

export function loadOptions(){
  Dicts({
    types: ['th_coupon_status','th_coupon_use_status'],
  }).then((res) => {
    options.value = res;
  });
  All({}).then((res) => {
    cateList.value = res.list;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'categoryId':
          item.componentProps.options = cateList.value;
          break;
      }
    }
  });
}


