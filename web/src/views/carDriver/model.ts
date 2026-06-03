import { h, ref } from 'vue';
import {NTag, SelectRenderLabel} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {defRangeShortcuts} from "@/utils/dateUtil";

export class State {
  public id = 0; // id
  public name = ''; // 真实姓名
  public nickname = '' // 昵称;
  public cooperateTypeId = null; // 合作类型ID
  public sex = 1; // 性别
  public phone = ''; // 手机号
  public phoneArea = ''; // 手机号区号
  public photo = ''; // 照片
  public age = null; // 年龄
  public workYears = null; // 从业年数
  public language = ''; // 语言能力
  public languageArr = [];
  public status = 2; // 1、启用 2、禁用
  public qualityMaterials = ''; //资质信息
  public qualityMaterialsArr = [];
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
  public deletedAt = ''; // deleted_at
  public settlementType = 1; // 结算类型
  public settlementRate = 0; // 结算比例
  public settlementId = null; // 结算模式

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
export const rules = {};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'keywords',
    component: 'NInput',
    label: '司机信息',
    componentProps: {
      placeholder: '请输入昵称或真实姓名或手机号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  // {
  //   field: 'workStatus',
  //   component: 'NSelect',
  //   label: '工作状态',
  //   defaultValue: null,
  //   componentProps: {
  //     placeholder: '请选择工作状态',
  //     options: [],
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  // {
  //   field: 'status',
  //   component: 'NSelect',
  //   label: '状态',
  //   defaultValue: 1,
  //   componentProps: {
  //     placeholder: '请选择状态',
  //     options: [],
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  {
    field: 'createAt',
    label: '加入时间',
    slot: 'createAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: '司机照片',
    key: 'photo',
    align: 'left',
    width: 100,
    render(row){
      return h(
        'img',
        {
          src: row.photo,
          loading: 'lazy',
          style: {
            width: '45px',
            // height: '45px',
            // borderRadius: '50%'
          },
        }
      )
    }
  },
  {
    title: '司机昵称',
    key: 'nickname',
    align: 'left',
    width: 120,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            {},
            {
              default: () => row.name,
            }
          ),
          h(
            NTag,
            {
              style: {
                marginRight: '6px',
                display:  row.isLeader==1 ? "inline-flex" : "none"
              },
              type: 'primary',
              bordered: false,
              size: 'small'
            },
            {
              default: () => row.isLeader==1 ? '车队长' : "普通司机",
            }
          )
        ]
      )

    }
  },
  {
    title: '真实姓名',
    key: 'name',
    align: 'left',
    width: 120,
  },
  {
    title: '手机号码',
    key: 'phone',
    align: 'left',
    width: 150,
    render(row){
      return row.phoneArea + '-' + row.phone
    }
  },
  {
    title: '绑定用户',
    key: 'memberId',
    align: 'left',
    width: 120,
    render(row) {
      if (row.memberId <= 0) {
        return `--`;
      }
      return h(
        'div',
        {},
        [
          h(
            'div',
            {},
            {
              default: () => row.memberDetail.memberNo
            }
          ),
          h(
            'div',
            {},
            {
              default: () => row.memberDetail.fullName
            }
          )
        ]
      )
    },
  },
  {
    title: '合作模式',
    key: 'cooperateTypeId',
    align: 'left',
    width: 80,
    render(row) {
      return row.cooperateTypeDetail.typeName
    },
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 80,
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
    title: '工作状态',
    key: 'workStatus',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.workStatus)) {
        return ``;
      }
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.driver_work_status, row.workStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.driver_work_status, row.workStatus),
        }
      );
    }
  },
  {
    title: '添加时间',
    key: 'createAt',
    align: 'left',
    width: 160,
  },
];

export const renderMemberLabel: SelectRenderLabel = (option) => {
  return option.memberNo + ' | ' + option.fullName + ' | ' + option.phone
};

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
  driver_work_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'driver_work_status'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        // case 'status':
        //   item.componentProps.options = options.value.sys_normal_disable;
        //   break;
        case 'workStatus':
          item.componentProps.options = options.value.driver_work_status;
          break;
      }
    }
  });
}


