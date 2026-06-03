import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {FormSchema} from "@/components/Form";
import {NTag} from 'naive-ui';

export class State {
  public id = 0; // id
  public terminalName = ''; // 终端名称
  public terminalType = 'VERIFY_PRINTER'; // 终端类型
  public brandModel = 'XPYUN503'; // 品牌型号
  public clientId = ''; // 开发者ID
  public clientSecret = ''; // 开发者秘钥
  public sn = ''; // 终端号
  public onlineStatus = ''; // // 在线状态
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
  terminalName: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入终端名称',
  },
  sn: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入终端编号',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'terminalName',
    component: 'NInput',
    label: '终端名称',
    componentProps: {
      placeholder: '请输入终端名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'sn',
    component: 'NInput',
    label: '终端号',
    componentProps: {
      placeholder: '请输入终端号',
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
    width: 80,
  },
  {
    title: '终端号',
    key: 'sn',
    align: 'left',
    width: 120,
  },
  {
    title: '终端名称',
    key: 'terminalName',
    align: 'left',
    width: 130,
  },
  {
    title: '终端类型',
    key: 'terminalType',
    align: 'left',
    width: 90,
    render(row) {
      if (isNullObject(row.terminalType)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.terminal_type, row.terminalType),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.terminal_type, row.terminalType),
        }
      );
    },
  },
  {
    title: '品牌型号',
    key: 'brandModel',
    align: 'left',
    width: 80,
    render(row) {
      if (isNullObject(row.brandModel)) {
        return ``;
      }
      if (row.terminalType === 'VERIFY_PRINTER') {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.verify_brand_model, row.brandModel),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.verify_brand_model, row.brandModel),
          }
        );
      }else if (row.terminalType === 'HAND_TERMINAL') {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.hand_brand_model, row.brandModel),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.hand_brand_model, row.brandModel),
          }
        );
      }else{
        return ``;
      }
    },
  },
  {
    title: '绑定门店',
    key: 'store_id',
    align: 'left',
    width: 130,
    render(row) {
      if (isNullObject(row.storeInfo)) {
        return ``;
      }
      return row.storeInfo.storeName
    }
  },
  {
    title: '绑定餐厅',
    key: 'restaurant_id',
    align: 'left',
    width: 130,
    render(row) {
      if (isNullObject(row.restaurantInfo)) {
        return ``;
      }
      return row.restaurantInfo.name
    }
  },
  {
    title: '创建时间',
    key: 'createAt',
    align: 'left',
    width: 150,
  },
  {
    title: '更新时间',
    key: 'updateAt',
    align: 'left',
    width: 150,
  },
];

// 字典数据选项
export const options = ref({
  terminal_type: [] as Option[],
  verify_brand_model: [] as Option[],
  hand_brand_model: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['terminal_type', 'verify_brand_model', 'hand_brand_model'],
  }).then((res) => {
    options.value = res;
    // for (const item of schemas.value) {
    //   switch (item.field) {
    //     case 'status':
    //       item.componentProps.options = options.value.sys_normal_disable;
    //       break;
    //   }
    // }
  });
}


