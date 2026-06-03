import {h, ref} from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { Dicts } from '@/api/dict/dict';
import { Option } from '@/utils/hotgo';
import {isNullObject} from "@/utils/is";
import {NTag} from "naive-ui";

export class State {
  public id = 0; // id
  public memberId = 0; // 用户ID
  public memberDetail = {
    fullName: ''
  };
  public memberNo = ''; // 会员号
  public phone = ''; // 手机
  public phoneArea = ''; // 手机区号
  public mail = ''; // 预定人邮箱
  public cancelReason = ''; // 注销原因
  public operatorId = ''; // 操作员ID
  public auditStatus = 'ALL'; // 审核状态
  public auditReason = ''; // 审核原因
  public auditTime = ''; // 审核时间
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public adminMemberUsername = ''; // 操作人名

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
    field: 'memberNo',
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
    field: 'createdAt',
    label: '申请时间',
    slot: 'createdAtSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: '会员号',
    key: 'memberNo',
    align: 'left',
    width: 200,
  },
 /* {
    title: '会员名',
    key: 'memberNo',
    align: 'left',
    width: 200,
    render(row) {
      if(row.memberDetail){
        return row.memberDetail.fullName
      }
      return '--'
    }
  },*/
  {
    title: '手机/邮箱',
    key: 'phone',
    align: 'left',
    width: 200,
    render(row) {
      return h(
        'div',
        [
          h(
            'div',
            [
              h(
                'span',
                {
                  class: 'c999 mr-1'
                },
                {
                  default: () => 'tel:',
                }
              ),
              h(
                'span',
                {},
                {
                  default: () => row.phone ? row.phoneArea + '-' + row.phone : '--',
                }
              ),
            ]
          ),
          h(
            'div',
            [
              h(
                'span',
                {
                  class: 'c999 mr-1'
                },
                {
                  default: () => 'mail:',
                }
              ),
              row.mail ? row.mail : '--',
            ]
          )
        ]
      )
    },
  },
  {
    title: '注销原因',
    key: 'cancelReason',
    align: 'left',
    width: 200,
  },
  {
    title: '审核状态',
    key: 'auditStatus',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.auditStatus)) {
        return ``;
      }
      if(row.auditStatus == 1){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'warning',
            bordered: false,
          },
          {
            default: () => '待审核',
          }
        );
      }else if(row.auditStatus == 2){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'success',
            bordered: false,
          },
          {
            default: () => '注销成功',
          }
        );
      }else if(row.auditStatus == 3){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'error',
            bordered: false,
          },
          {
            default: () => '审核拒绝',
          }
        );
      }
      return ""
    },
  },
  {
    title: '申请时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },
];

// 字典数据选项
export const options = ref({
  order_status: [] as Option[],
  booking_status: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  Dicts({
    types: ['order_status','booking_status'],
  }).then((res) => {
    options.value = res;
  });
}


