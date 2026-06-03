import { ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Option } from '@/utils/hotgo';
import { Dicts } from '@/api/dict/dict';

export interface ScopeItem {
  productId: number;
  skuId: number;
}

export interface ScopeSkuOption {
  id: number;
  productId: number;
  name: string;
}

export interface ScopeOption {
  id: number;
  title: string;
  skuList: ScopeSkuOption[];
}

export class State {
  public id = 0;
  public name = '';
  public mobile = '';
  public username = '';
  public password = '';
  public status = 1;
  public statusName = '';
  public createdAt = '';
  public updatedAt = '';
  public isAllScope = true;
  public scopeItems: ScopeItem[] = [];
  public scopeSelectedKeys: Array<string | number> = [];

  constructor(state?: Partial<State>) {
    if (state) Object.assign(this, state);
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) return cloneDeep(state);
    return new State(state);
  }
  return new State();
}

export const rules = {
  name: { required: true, trigger: ['blur', 'input'], message: '请输入姓名' },
  mobile: { required: true, trigger: ['blur', 'input'], message: '请输入电话' },
  username: { required: true, trigger: ['blur', 'input'], message: '请输入登录账号' },
  status: { required: true, trigger: ['blur', 'change'], type: 'number', message: '请选择状态' },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '姓名',
    componentProps: { placeholder: '请输入姓名' },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    componentProps: {
      placeholder: '请选择状态',
      options: [
        { label: '全部', value: 0 },
        { label: '启用', value: 1 },
        { label: '禁用', value: 2 },
      ],
    },
  },
]);

export const options = ref({
  sys_normal_disable: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['sys_normal_disable', 'driver_work_status'],
  }).then((res) => {
    options.value = res;
  });
}
