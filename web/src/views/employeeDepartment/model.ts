import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { getOptionLabel, getOptionTag } from '@/utils/hotgo';
import { isNullObject } from '@/utils/is';

export class State {
  public id = 0; // 部门ID
  public name = ''; // 部门名称
  public parentId = null; // 上级部门ID
  public level = 1; // 部门层级
  public path = ''; // 部门路径
  public managerId = null; // 部门负责人ID
  public description = ''; // 部门描述
  public sort = 0; // 排序
  public status = 1; // 状态
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 修改时间

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

export const options = ref<any>({
  sys_normal_disable: [],
  department_tree: [],
});

export const schemas: FormSchema[] = [
  {
    field: 'name',
    component: 'NInput',
    label: '部门名称',
    componentProps: {
      placeholder: '请输入部门名称',
      onKeyup: (e: KeyboardEvent) => {
        if (e.key === 'Enter') {
          e.preventDefault();
        }
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    componentProps: {
      placeholder: '请选择状态',
      options: [
        {
          label: '正常',
          value: 1,
        },
        {
          label: '禁用',
          value: 2,
        },
      ],
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        // 处理时间范围
      },
    },
  },
];

export const columns = [
  {
    title: '部门名称',
    key: 'name',
    width: 200,
    tree: true,
    render(row) {
      return h('span', { style: { fontWeight: 'bold' } }, row.name);
    },
  },
  {
    title: '部门层级',
    key: 'level',
    width: 80,
    render(row) {
      const levelColors = ['', 'primary', 'info', 'success', 'warning', 'error'];
      const color = levelColors[row.level] || 'default';
      return h(NTag, { type: color, size: 'small' }, { default: () => `${row.level}级` });
    },
  },
  {
    title: '员工数量',
    key: 'employeeCount',
    width: 100,
    render(row) {
      return h(NTag, { type: 'info', size: 'small' }, { default: () => `${row.employeeCount || 0}人` });
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
    title: '创建时间',
    key: 'createdAt',
    width: 180,
  },
];

export async function loadOptions() {
  options.value = await Dicts({
    types: ['sys_normal_disable'],
  });
}
