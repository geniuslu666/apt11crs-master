import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public typeId = 1; // 地点类型id
  public name = ''; // 地点名称（后台）
  public subName = ''; // 地点名称（app-多语）
  public nameLanguage = null;
  public airportCode = ''; // 机场代码
  public terminalName = ''; // 航站楼
  public detailAddress = ''; // 详细地址-多语
  public addressLanguage = null;
  public ggLat = '34.67100087743556'; // 谷歌纬度
  public ggLng = '135.49982492658245'; // 谷歌经度
  public lat = ''; // 纬度
  public lng = ''; // 经度
  public status = 1; // 状态1、启用 2、禁用
  public propertyId = 1; // 物业id
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
  typeId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入地点类型id',
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入地点名称（后台）',
  },
  subName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入地点名称（app-多语）',
  },
  airportCode: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入机场代码',
  },
  terminalName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入航站楼',
  },
  detailAddress: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入详细地址-多语',
  },
  ggLat: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入谷歌纬度',
  },
  ggLng: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入谷歌经度',
  },
  lat: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入纬度',
  },
  lng: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入经度',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入状态',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '地点名称',
    componentProps: {
      placeholder: '请输入地点名称（后台）',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态',
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
    title: 'id',
    key: 'id',
    align: 'left',
    width: 100,
  },
  {
    title: '地点类型',
    key: 'typeId',
    align: 'left',
    width: 110,
    render(row){
      return row.addressTypeDetail.typeName
    }
  },
  {
    title: '地点名称',
    key: 'name',
    align: 'left',
    width: 150,
  },
  {
    title: '地点名称（app）',
    key: 'subName',
    align: 'left',
    width: 150,
  },
  /*{
    title: '机场代码',
    key: 'airportCode',
    align: 'left',
    width: -1,
  },*/
  {
    title: '详细地址',
    key: 'detailAddress',
    align: 'left',
    width: -1,
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
    key: 'createAt',
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
    for (const item of schemas.value) {
      switch (item.field) {
        case 'status':
          item.componentProps.options = options.value.sys_normal_disable;
          break;
      }
    }
  });
}


