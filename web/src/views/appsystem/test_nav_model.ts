import { h, ref } from 'vue';
import {NImage, NInput, NSwitch} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import {Option, errorImg} from '@/utils/hotgo';
import {Switch, Sort} from "@/api/pmsIndexNav";

const $message = window['$message'];

export class State {
  public id = 0; // id
  public name = ''; // 门店名称
  public image = ''; // 图集
  public appLink = ''; // app链接
  public wxLink = ''; // wx链接
  public sort = 0; // 排序(越大越靠前)
  public status = 1; // 状态
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

// 表格搜索表单

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '名称',
    key: 'name',
    align: 'left',
    width: 100,
  },
  {
    title: '图标',
    key: 'image',
    align: 'left',
    width: 100,
    render(row) {
      return h(NImage, {
        width: 32,
        height: 32,
        src: row.image,
        onError: errorImg,
        style: {
          width: '32px',
          height: '32px',
          'max-width': '100%',
          'max-height': '100%',
        },
      });
    },
  },
  {
    title: '创建时间',
    key: 'createAt',
    align: 'left',
    width: 180,
  },
  {
    title: '更新时间',
    key: 'updateAt',
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


