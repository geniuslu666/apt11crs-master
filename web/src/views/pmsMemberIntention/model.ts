import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defShortcuts } from '@/utils/dateUtil';
import {getDeptOption} from "@/api/org/dept";
import {All} from "@/api/pmsProperty";
import {getRoleOption} from "@/api/system/role";
import {Dicts} from "@/api/dict/dict";
import {isNullObject} from "@/utils/is";
import {NButton, NTag, NPopover} from "naive-ui";
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import defaultImg from "@/assets/images/mrtx.png";
import {QuestionCircleOutlined} from "@vicons/antd";

export class State {
  public id = 0; // 主键ID
  public memberId = 0; // 意向用户ID
  public name = ''; // 称呼
  public sex = 1; // 1、男 2、女
  public country = ''; // 国家/地区
  public phone = ''; // 手机号
  public phoneArea = ''; // 手机区号
  public language = '';// 语言喜好
  public mail = ''; // 邮箱
  public wechatNo = ''; // 微信号
  public lineId = ''; // LINE ID
  public remark = ''; // 备注
  public createAt = ''; // 创建时间
  public updateAt = ''; // 修改时间
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
  memberId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入意向用户ID',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '称呼',
    componentProps: {
      placeholder: '请输入称呼',
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
    field: 'mail',
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
    field: 'createAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetime',
      clearable: true,
      shortcuts: defShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [

  {
    title: '会员信息',
    key: 'memberId',
    align: 'left',
    width: 200,
    render(row) {
      if(!row.member){
        return "--"
      }
      return h(
        'div',
        {
          style: {
            display: 'flex',
            alignItems: 'center'
          }
        },
        [
          h(
            'img',
            {
              src: row.member.avatar ? row.member.avatar : defaultImg,
              style: {
                width: '38px',
                height: '38px',
                borderRadius: '50%'
              },
              onError: (e) => {
                e.target.src = defaultImg
              }
            }
          ),
          h(
            'div',
            {
              class:'flex-item',
              style: {
                paddingLeft: '8px',
                lineHeight: '20px'
              }
            },
            [
              h(
                'div',
                {
                  style:{
                    color: '#1664FF',
                    fontSize: '14px'
                  }
                },
                {
                  default: () => row.member.memberNo,
                }
              ),
              h(
                'div',
                {
                  style:{
                    color: '#3D3D3D',
                    fontSize: '14px'
                  }
                },
                {
                  default: () => row.member.fullName,
                }
              ),
            ]
          )
        ]
      )
    }
  },
  // {
  //   title: '称呼',
  //   key: 'name',
  //   align: 'left',
  //   width: 100,
  // },
  // {
  //   title: '性别',
  //   key: 'sex',
  //   align: 'left',
  //   width: 90,
  //   render(row) {
  //     if (isNullObject(row.sex)) {
  //       return ``;
  //     }
  //     return h(
  //       NTag,
  //       {
  //         style: {
  //           marginRight: '6px',
  //         },
  //         type: getOptionTag(options.value.sys_user_sex, row.sex),
  //         bordered: false,
  //       },
  //       {
  //         default: () => getOptionLabel(options.value.sys_user_sex, row.sex),
  //       }
  //     );
  //   },
  // },
  {
    title: '国家/地区',
    key: 'country',
    align: 'left',
    width: 90,
  },
  {
    title: '语言喜好',
    key: 'language',
    align: 'left',
    width: 90,
    render(record) {
      if (isNullObject(record.language)) {
        return ``;
      }
      if(record.language == '简体中文'){
        return h(
          'div',
          {
            style: {
              color: '#F56C6C',
              padding: '0 5px',
              height: '22px',
              background: '#FEF0F0',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '简体中文',
          }
        );
      }else if(record.language == 'English'){
        return h(
          'div',
          {
            style: {
              color: '#EFA020',
              padding: '0 5px',
              height: '22px',
              background: '#FDF6EC',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => 'English',
          }
        );
      }else if(record.language == '日本語'){
        return h(
          'div',
          {
            style: {
              color: '#3F9EFF',
              padding: '0 5px',
              height: '22px',
              background: '#ECF5FF',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '日本語',
          }
        );
      }else if(record.language == '한국어'){
        return h(
          'div',
          {
            style: {
              color: '#26A763',
              padding: '0 5px',
              height: '22px',
              background: '#E3F4EB',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '한국어',
          }
        );
      }else if(record.language == '繁体中文'){
        return h(
          'div',
          {
            style: {
              color: '#434447',
              padding: '0 5px',
              height: '22px',
              background: '#EEEEEE',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '繁体中文',
          }
        );
      }
    },
  },
  {
    title: '手机号',
    key: 'phone',
    align: 'left',
    width: 120,
    render(row) {
      return row.phoneArea+ " "+ row.phone
    },
  },

  {
    title: '邮箱',
    key: 'mail',
    align: 'left',
    width: 120,
  },
  {
    title: '留言备注',
    key: 'remark',
    align: 'left',
    width: 120,
    render(row) {
      return h(
        NPopover,
        null,
        {
          trigger:()=>
            h(
              NButton,
              {
                text: true,
                type: 'primary'
              },
              {
                default: () => '查看',
              }
            ),
          default: () => row.remark,
        },
      )
    },
  },
  {
    title: '提交时间',
    key: 'createAt',
    align: 'left',
    width: 155,
    ellipsis: false,
  },
];

export const options = ref<any>({
  sys_user_sex: [],
  language: [] as Option[],
});

export async function loadOptions() {
  const tmpOptions = await Dicts({
    types: ['sys_user_sex', 'language'],
  });
  options.value.sys_user_sex = tmpOptions?.sys_user_sex;
  options.value.language = tmpOptions?.language;
  // console.log("options.value", options.value)
}

