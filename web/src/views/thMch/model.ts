import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import {NSwitch, NTag} from "naive-ui";
import {Switch} from "@/api/thMch";
import {All} from "@/api/thMchCategory";
import {isNullObject} from "@/utils/is";
import {Option, getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";

const $message = window['$message'];

export class State {
  public id = 0; // id
  public categoryId = null; // 分类ID
  public thMchCategoryName = ''; // 分类名称
  public name = ''; // 名称
  public logo = ''; // logo
  public contactInfo = ''; // 联系信息
  public sort = 0; // 排序
  public status = 1; // 状态
  public storeOnNum = 0; //启用中门店
  public storeOffNum = 0; // 禁用中门店
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at
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

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '商户名称',
    componentProps: {
      placeholder: '请输入商户名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'categoryId',
    component: 'NSelect',
    label: '商户分类',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择商户分类',
      options: [],
      labelField: 'name',
      valueField: 'id',
      onUpdateValue: (e: any) => {
        console.log(typeof(e));
      },
    },
  },
  {
    field: 'contactInfo',
    component: 'NInput',
    label: '联系信息',
    componentProps: {
      placeholder: '请输入联系信息',
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
      options: [
        {
          labelField: '正常',
          valueField: 1,
        },
        {
          labelField: '停用',
          valueField: 2,
        }
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(typeof(e));
      },
    },
  },
]);

// 表格列

export const cateList = ref([]);
// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
});

export function loadOptions(){
  Dicts({
    types: ['sys_normal_disable'],
  }).then((res) => {
    options.value = res;
  });
  All({}).then((res) => {
    cateList.value = res.list;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'categoryId':
          item.componentProps.options = cateList.value;
          break;
      }
    }
  });
}

export function loadDictOptions(){
  Dicts({
    types: ['sys_normal_disable'],
  }).then((res) => {
    options.value = res;
  });
}
