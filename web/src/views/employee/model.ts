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
import {GetDepartmentTree} from "@/api/employeeDepartment";

export class State {
  public id = 0; // id
  public name = ''; // 员工姓名
  public phone_area = ''; // 区号
  public phone = ''; // 手机号
  public memberId = null; // 绑定用户ID
  public departmentId = null; // 员工部门ID
  public employeeNo = ''; // 员工编号
  public status = 1; // 状态   1、 开启   2、禁用
  public remark = ''; // 备注
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

// 表单验证规则
export const rules = {
  phone: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请填写手机号',
    // validator: validate.phone,
  },
  email: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
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
    field: 'departmentId',
    component: 'NTreeSelect',
    label: '部门',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择部门',
      options: [],
      labelField: 'name',
      valueField: 'id',
      childrenField: 'children',
      clearable: true,
      filterable: true,
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
    render(row) {
      if (isNullObject(row.departmentDetail) || row.departmentId <= 0) {
        return ``;
      }
      return row.departmentDetail.name;
    }
  },
  {
    title: '手机号',
    key: 'phone',
    align: 'left',
    width: 160,
    render(row) {
      if (isNullObject(row.phone)) {
        return ``;
      }
      return row.phoneArea + ' ' + row.phone;
    },
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
    title: '会员',
    key: 'memberId',
    align: 'left',
    width: -1,
    render(row) {
      if (isNullObject(row.memberId) || row.memberId <= 0) {
        return `--`;
      }

      var memberDetail = row.memberDetail
      if(!memberDetail){
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
              default: () => row.memberDetail.fullName
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
              default: () => "编号："+row.memberDetail.memberNo+"，名称："+row.memberDetail.fullName+"，手机："+row.memberDetail.phoneArea + "-" + row.memberDetail.phone +"，邮箱：" + row.memberDetail.mail,
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
  GetDepartmentTree({}).then((res) => {
    for (const item of schemas.value) {
      switch (item.field) {
        case 'departmentId':
          item.componentProps.options = res.list;
          break;
      }
    }
  });
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


