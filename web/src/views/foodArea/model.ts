import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public pid = 0; // 上级ID
  public level = 0; // 区域级别
  public tree = ''; // 区域名称
  public areaName = ''; // 区域名称
  public areaStatus = 1; // 1、启用 2、禁用
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
  areaName: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入区域名称',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'id',
    componentProps: {
      placeholder: '请输入id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'areaStatus',
    component: 'NSelect',
    label: '1、启用 2、禁用',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择1、启用 2、禁用',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '区域名称',
    key: 'areaName',
    align: 'left',
    width: -1,
  },
  {
    title: '状态',
    key: 'areaStatus',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.areaStatus)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.areaStatus),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.areaStatus),
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
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'areaStatus':
          item.componentProps.options = options.value.sys_normal_disable;
          break;
      }
    }
  });
}


