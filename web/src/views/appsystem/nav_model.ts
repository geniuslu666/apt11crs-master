import { h, ref } from 'vue';
import {NImage, NInput, NSwitch, NTag} from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import {Option, errorImg, getOptionTag, getOptionLabel} from '@/utils/hotgo';
import {Switch, Sort, MinappStatus} from "@/api/pmsIndexNav";
import {isNullObject} from "@/utils/is";

const $message = window['$message'];

export class State {
  public id = 0; // id
  public name = ''; // 门店名称
  public nameLanguage = null;
  public tag = ''; // 标签
  public tagLanguage = null;
  public image = ''; // 图集
  public appLink = ''; // app链接
  public wxLink = ''; // wx链接
  public sort = 0; // 排序(越大越靠前)
  public status = 1; // 状态
  public minappStatus = 1; // 小程序是否显示 1 显示 2 不显示
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间
  public deletedAt = ''; // deleted_at
  public chain = "IN"; // 内外联
  public linkOpenType = 1; // 外链跳转方式 1-内部webview 2-外部浏览器

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
    title: '标签',
    key: 'tag',
    align: 'left',
    width: 100,
    render(row){
      return row.tag ? row.tag : '--'
    }
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
    title: '是否启用',
    key: 'status',
    align: 'left',
    width: 100,
    render(row) {
      return h(NSwitch, {
        value: row.status === 1,
        checked: '开启',
        unchecked: '关闭',
        onUpdateValue: function (e) {
          row.status = e ? 1 : 2;
          Switch({ id: row.id, key: 'switch', status: row.status }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
  {
    title: '小程序显示',
    key: 'minappStatus',
    align: 'left',
    width: 100,
    render(row) {
      return h(NSwitch, {
        value: row.minappStatus === 1,
        checked: '显示',
        unchecked: '不显示',
        onUpdateValue: function (e) {
          row.minappStatus = e ? 1 : 2;
          MinappStatus({ id: row.id, key: 'switch', status: row.minappStatus }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
  {
    title: '内外联',
    key: 'chain',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.chain)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.chain, row.chain),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.chain, row.chain),
        }
      );
    },
  },
  {
    title: '排序',
    key: 'sort',
    width: 80,
    render(row) {
      return h(NInput, {
        value: row.sort,
        onUpdateValue(v) {
          Sort({ id: row.id, sort: v }).then((_res) => {
            $message.success('操作成功');
          });
          row.sort = v
        }
      })
    }
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
  chain: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'chain'],
  }).then((res) => {
    options.value = res;
  });
}


