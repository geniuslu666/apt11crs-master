import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; //
  public restaurantId = null; // 餐厅ID
  public title = ''; // 通知标题
  public content = ''; // 公告内容
  public sort = 0; // 排序
  public status = 1; // 状态
  public needUserConfirm = 1; // 用户确认
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
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
  restaurantId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入餐厅ID',
  },
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入通知标题',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入公告内容',
  },
  sort: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入排序',
  },
};

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '通知标题',
    key: 'title',
    align: 'left',
    width: -1,
  },
  {
    title: '排序',
    key: 'sort',
    align: 'left',
    width: -1,
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
    title: '用户确认',
    key: 'needUserConfirm',
    align: 'left',
    width: -1,
    render(row) {
      if(row.needUserConfirm == 1){
        return h(
          NTag,
          {
            style: {

            },
            type: 'warning',
            bordered: false,
          },
          {
            default: () => "需要用户确认",
          }
        );
      }else if(row.needUserConfirm == 2){
        return h(
          NTag,
          {
            style: {

            },
            type: 'info',
            bordered: false,
          },
          {
            default: () => "不需要确认",
          }
        );
      }

    }
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
  });
}


