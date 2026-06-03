import { h, ref } from 'vue';
import {NImage, NTag} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import defaultImg from "@/assets/images/onerror.png";
import actEmptyImg from "@/assets/images/emp_act_df.png";

export interface LanguageObject {
  en: string;
  zh: string;
  ja: string;
  ko: string;
  zh_CN: string;
}

export interface ThCoupon {
  couponId: number;
  availableQuantity: number;
  perDayAvailable: number;
  perDayVerify: number;
}

export class State {
  public id = 0; // id
  public name = ''; // 活动名称
  public nameLanguage: LanguageObject | null = null;
  public cover = null; // 活动封面
  public description = ''; // 活动描述
  public descriptionLanguage: LanguageObject | null = null;
  public validityType = 1; // 有效期类型
  public startTime = null; // 活动开始时间
  public endTime = null; // 活动结束时间
  public status = 1; // 状态：1-未开始 2-进行中 3-已结束
  public manualClosed = 0; // 是否手动关闭
  public couponValidity = 1;// 券有效期
  public restrictionType = 1; // 限制类型
  public rule = ''; // 活动规则
  public sort = 0; // 排序
  public isEnabled = 1; // 是否启用
  public thCouponsArr: ThCoupon[] = []; // 礼品券Arr
  public employeeIds = [] as number[]; // 员工Ids
  public departmentIds = [] as number[]; // 部门Ids
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public deletedAt = ''; // 删除时间
  public limitWeek = ''; // 周末领取限制 6-周六不允许,7-周日不允许，多个用逗号分隔

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

};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '活动名称',
    componentProps: {
      placeholder: '请输入活动名称',
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '活动名称',
    key: 'name',
    align: 'left',
    width: 150,
  },
  {
    title: '描述',
    key: 'description',
    align: 'left',
    width: 250,
  },
  {
    title: '图片',
    key: 'cover',
    align: 'left',
    width: 80,
    render(row) {

      var pic = row.cover;

      if(!pic){
        return h(
          NImage,
          {
            style: {
              width: '45px',
              // height: '45px',
              // borderRadius: '50%'
            },
            src: actEmptyImg,
            onError: (e) => {
              e.target.src = defaultImg
            }
          },

        )
      }else{
        return h(
          'div',
          {
            class: 'flex-row',
          },
          [
            h(
              NImage,
              {
                src: pic ? pic : defaultImg,
                loading: 'lazy',
                style: {
                  width: '45px',
                  height: '45px',
                  // borderRadius: '50%'
                },
                onError: (e) => {
                  e.target.src = defaultImg
                }
              }
            ),
          ]
        );
      }
    },
  },
  {
    title: '是否启用',
    key: 'is_enabled',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.isEnabled)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.isEnabled),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.isEnabled),
        }
      );
    },
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 80,
    render(row){
      if(row.status == 1){
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
            default: () => '未开始',
          }
        );
      }
      if(row.status == 2){
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
  },
  {
    title: '有效期',
    key: 'startTime',
    align: 'left',
    width: 150,
    render(row){
      if (row.validityType == 2){
        return '长期有效'
      }
      return row.startDate + ' ~ ' + row.endDate
    }
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },
];

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable'],
  }).then((res) => {
    options.value = res;
  });
}


