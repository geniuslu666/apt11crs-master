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
  public status = 1; // 1、启用 2、禁用
  public workStatus = 'WORKING'; //
  public memberId = 0;
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
  public deletedAt = ''; // deleted_at
  public settlementType = 1; // 结算类型
  public settlementRate = 0; // 结算比例
  public settlementId = null; // 结算模式
  public settlementObject = 'ISP';

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
    label: '服务商信息',
    componentProps: {
      placeholder: '请输入名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'workStatus',
    component: 'NSelect',
    label: '工作状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择工作状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
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
    label: '创建时间',
    slot: 'createAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: '名称',
    key: 'name',
    align: 'left',
    width: 120,
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
    title: '营业状态',
    key: 'workStatus',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.workStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.isp_work_status, row.workStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.isp_work_status, row.workStatus),
        }
      );
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
    title: '添加时间',
    key: 'createAt',
    align: 'left',
    width: 180,
  },
];

export const renderMemberLabel: SelectRenderLabel = (option) => {
  return option.memberNo + ' | ' + option.fullName + ' | ' + option.phone
};

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
  isp_work_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'isp_work_status'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        // case 'status':
        //   item.componentProps.options = options.value.sys_normal_disable;
        //   break;
        case 'workStatus':
          item.componentProps.options = options.value.isp_work_status;
          break;
      }
    }
  });
}


