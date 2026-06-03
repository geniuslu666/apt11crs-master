import { h, ref } from 'vue';
import { NTag, SelectRenderLabel, NTooltip, NIcon } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import {QuestionCircleOutlined} from '@vicons/antd';

export class State {
  public id = 0; // id
  public name = ''; // 员工姓名
  public department = ''; // 员工部门
  public phone = ''; // 手机号
  public email = ''; // 邮箱
  public rate = ''; // 返佣比例
  public status = 1; // 状态   1、 开启   2、禁用
  public remark = ''; // 备注
  public balance = null; // 可提现账户
  public allBalance = null; // 总账户
  public minWithdrawalAmount = null; // 最低可提现额
  public serviceCharge = null; // 手续费
  public afterDay = 0; // 预计提现时间周期
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at
  public deletedAt = ''; // deleted_at
  public pmsMemberId = null;

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
  phone: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.phone,
  },
  email: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
  },
  rate: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'number',
    validator: validate.percentage,
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请选择状态',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '员工姓名',
    componentProps: {
      placeholder: '请输入员工姓名',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'department',
    component: 'NInput',
    label: '员工部门',
    componentProps: {
      width: '100px',
      placeholder: '请输入员工部门',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'phone',
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
    field: 'email',
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
    field: 'rate',
    component: 'NInput',
    label: '返佣比例',
    componentProps: {
      placeholder: '请输入返佣比例',
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
  {
    field: 'createdAt',
    label: '创建时间',
    slot: 'createdAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: 80,
  },
  {
    title: '员工姓名',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '员工部门',
    key: 'department',
    align: 'left',
    width: -1,
  },
  {
    title: '手机号',
    key: 'phone',
    align: 'left',
    width: 130,
  },
  {
    title: '邮箱',
    key: 'email',
    align: 'left',
    width: 170,
  },
  {
    title: '返佣比例%',
    key: 'rate',
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
    title: '可提现账户',
    key: 'balance',
    align: 'left',
    width: -1,
  },
  {
    title: '总账户',
    key: 'allBalance',
    align: 'left',
    width: -1,
  },
  {
    title: '会员',
    key: 'pmsMemberFullName',
    align: 'left',
    width: -1,
    render(row) {
      var mail = row.pmsMemberMail
      if(!mail){
        mail = '无'
      }
      if (row.pmsMemberId <= 0) {
        return `--`;
      }
      return h(
        'div',
        {
          style: {
            display: 'flex',
            alignItems: 'center'
          },
        },
        [
          h(
            'span',
            {},
            {
              default: () => row.pmsMemberFullName
            }
          ),
          h(
            NTooltip,
            null,
            {
              trigger:()=>
                h(
                  NIcon,
                  {
                    size: 20,
                    style: {
                      marginLeft: '5px',
                    },
                  },
                  {
                    default: () => h(QuestionCircleOutlined),
                  }
                ),
              default: () => "编号："+row.pmsMemberMemberNo+"，名称："+row.pmsMemberFullName+"，手机："+row.pmsMemberPhone+"，邮箱：" + mail,
            },
          )
        ],
      );
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 160,
  },
];

export const renderMemberLabel: SelectRenderLabel = (option) => {
  return option.memberNo + ' | ' + option.fullName + ' | ' + option.phone
};

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


