import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defShortcuts, defRangeShortcuts } from '@/utils/dateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

export class State {
  public id = 0; // id
  public memberId = 0; // 会员ID
  public loginTime = ''; // 登录时间
  public loginType = null; // 登录方式
  public loginIp = ''; // 登录IP
  public token = ''; // 登录token
  public expirTime = ''; // 过期时间
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at

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
export const schemas = ref<FormSchema[]>([

  {
    field: 'loginType',
    component: 'NSelect',
    label: '登录方式',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择登录方式',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },

  {
    field: 'pmsMemberMemberNo',
    component: 'NInput',
    label: '会员号',
    componentProps: {
      placeholder: '请输入会员号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'pmsMemberFullName',
    component: 'NInput',
    label: '全名',
    componentProps: {
      placeholder: '请输入全名',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'pmsMemberPhone',
    component: 'NInput',
    label: '手机号',
    componentProps: {
      placeholder: '请输入手机号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'pmsMemberPhoneArea',
    component: 'NInput',
    label: '手机区号',
    componentProps: {
      placeholder: '请输入手机区号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'pmsMemberMail',
    component: 'NInput',
    label: '邮箱',
    componentProps: {
      placeholder: '请输入邮箱',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'pmsMemberSource',
    component: 'NInput',
    label: '注册来源',
    componentProps: {
      placeholder: '请输入注册来源',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [

  {
    title: '会员ID',
    key: 'memberId',
    align: 'left',
    width: 80,
  },

  {
    title: '登录方式',
    key: 'loginType',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.loginType)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.login_type, row.loginType),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.login_type, row.loginType),
        }
      );
    },
  },
  {
    title: '登录IP',
    key: 'loginIp',
    align: 'left',
    width: 150,
  },
  {
    title: '登录时间',
    key: 'loginTime',
    align: 'left',
    width: 200,
  },
  {
    title: '过期时间',
    key: 'expirTime',
    align: 'left',
    width: 200,
  },
  {
    title: '会员号',
    key: 'pmsMemberMemberNo',
    align: 'left',
    width: 100,
  },
  // {
  //   title: '头像',
  //   key: 'pmsMemberAvatar',
  //   align: 'left',
  //   width: -1,
  // },
  {
    title: '全名',
    key: 'pmsMemberFullName',
    align: 'left',
    width: 200,
  },
  {
    title: '手机区号',
    key: 'pmsMemberPhoneArea',
    align: 'left',
    width: 80,
  },
  {
    title: '手机号',
    key: 'pmsMemberPhone',
    align: 'left',
    width: 150,
  },
  {
    title: '邮箱',
    key: 'pmsMemberMail',
    align: 'left',
    width: 180,
  },
  {
    title: '注册来源',
    key: 'pmsMemberSource',
    align: 'left',
    width: 100,
  },
  {
    title: '设备号',
    key: 'mdCode',
    align: 'left',
    width: 150,
    resizable: true,
  },
  {
    title: '手机型号',
    key: 'mpModel',
    align: 'left',
    width: 120,
  },
];

// 字典数据选项
export const options = ref({
  login_type: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['login_type'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'loginType':
          item.componentProps.options = options.value.login_type;
          break;
      }
    }
  });
}
