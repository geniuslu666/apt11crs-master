import {h, ref} from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import {defRangeShortcuts} from '@/utils/dateUtil';
import {All} from "@/api/pmsMemberLevel/index";
import {getConfig} from "@/api/sys/config";
import defaultImg from "@/assets/images/mrtx.png";
import {isNullObject} from "@/utils/is";
import {NInput, NTag, NTooltip} from "naive-ui";
import {getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {options} from "@/views/foodCooperateType/model";
import {Dicts} from "@/api/dict/dict";

export class State {
  public id = 0; // 主键
  public memberNo = ''; // 会员号
  public firstName = ''; // 名
  public lastName = ''; // 姓
  public fullName = ''; // 全名
  public level = 0; // 等级
  public groupId = 0; // 会员分组
  public phone = ''; // 手机号
  public mail = ''; // 邮箱
  public birthday = ''; // 生日
  public balance = 0; // 积分
  public lastReferrer = 0; // 最后推荐人
  public referrer = 0; // 推荐人
  public source = ''; // 注册来源  IOS,Andriod,h5
  public lastLoginIp = ''; // 上次登录IP
  public lastLogin = ''; // 上次登录时间
  public registerIp = ''; // 注册IP
  public registerTime = ''; // 注册时间
  public address = ''; // 地址
  public nationality = ''; // 国籍
  public createAt = ''; // 创建时间
  public updateAt = ''; // 更新时间

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
  firstName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入名',
  },
  lastName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入姓',
  },
};

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
    field: 'fullName',
    component: 'NInput',
    label: '昵称',
    componentProps: {
      placeholder: '请输入会员昵称',
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
    field: 'level',
    component: 'NSelect',
    label: '会员等级',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择会员等级',
      options: [],
      labelField: 'levelName',
      valueField: 'id',
      onUpdateValue: (e: any) => {
        console.log(typeof(e));
      },
    },
  },
  {
    field: 'createdAt',
    label: '创建时间',
    slot: 'createdAtSlot',
  },
  {
    field: 'lastLogin',
    label: '登录时间',
    slot: 'lastLoginSlot',
  },
]);

// 表格列
export const columns = [
  {
    title: '会员信息',
    key: 'memberNo',
    align: 'left',
    width: 200,
    resizable: true,
    render(row) {
      return h(
        'div',
        {
          class: 'flex-row',
        },
        [
          h(
            'img',
            {
              src: row.avatar ? row.avatar : defaultImg,
              loading: 'lazy',
              style: {
                width: '45px',
                height: '45px',
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
                paddingLeft: '10px',
                lineHeight: '22px'
              }
            },
            [
              h(
                'div',
                [
                  h(
                    'span',
                    {
                      class: 'cblue f14 fw'
                    },
                    {
                      default: () => row.memberNo,
                    }
                  ),
                  // h(
                  //   'a',
                  //   {
                  //     class: 'c999 f12 ml-1',
                  //     directives: [{
                  //       name: 'copy',
                  //       value: 'red'
                  //     }],
                  //   },
                  //   {
                  //     default: () => '复制',
                  //   }
                  // )
                ]
              ),
              h(
                'div',
                {},
                {
                  default: () => row.lastName + " " + row.firstName,
                }
              )
            ]
          )
        ]
      );
    },
  },
  {
    title: '会员分组',
    key: 'groupId',
    align: 'left',
    width: 150,
    render(row){
      if(row.groupId > 0 && row.memberGroup){
        return row.memberGroup.memberGroup
      }
      return '--'

    }
  },
  {
    title: '会员等级',
    key: 'level',
    align: 'left',
    width: 150,
    render(row){
      return row.memberLevel.levelName
    }
  },
  {
    title: '推荐信息',
    key: 'phoneArea',
    align: 'left',
    width: 150,
    render(row){
      if(row.referrer <= 0 && row.lastReferrer <= 0){
        return '无'
      }
      if(recommendModel.value == 'FIRST'){
        if(!row.referrerDetail){
          return '无'
        }else{
          if(row.referrerDetail.rebateMode == 'CHANNEL'){
            return '渠道：' + row.referrerDetail.memberNo
          }else if(row.referrerDetail.rebateMode == 'STAFF'){
            return '员工：' + row.referrerDetail.memberNo
          }else{
            return '会员：' + row.referrerDetail.memberNo
          }
        }
      }else{
        if(!row.LastReferrerDetail){
          return '无'
        }else{
          if(row.LastReferrerDetail.rebateMode == 'CHANNEL'){
            return '渠道：' + row.LastReferrerDetail.memberNo
          }else if(row.LastReferrerDetail.rebateMode == 'STAFF'){
            return '员工：' + row.LastReferrerDetail.memberNo
          }else{
            return '会员：' + row.LastReferrerDetail.memberNo
          }
        }
      }
    }
  },
  {
    title: '手机/邮箱',
    key: 'phone',
    align: 'left',
    width: 200,
    render(row) {
      let htmlStr = h(
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
              // row.phone ? h(
              //   'a',
              //   {
              //     class: 'c999 f12 ml-1',
              //   },
              //   {
              //     default: () => '复制',
              //   }
              // ) : ''
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
              // row.mail ? h(
              //   'a',
              //   {
              //     class: 'c999 f12 ml-1',
              //   },
              //   {
              //     default: () => '复制',
              //   }
              // ) : ''
            ]
          )
        ]
      );
      return h(
        NTooltip,
        null,
        {
          trigger:()=>
            htmlStr,
          default: () => htmlStr,
        },
      )
    },
  },

  {
    title: '会员积分/经验',
    key: 'exp',
    align: 'left',
    sorter: true, // 单列排序
    width: 130,
    render(row){
      return row.balance + '/' + row.exp
    }
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 100,
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
    title: '注册来源',
    key: 'source',
    align: 'left',
    width: 100,
  },
  {
    title: '注册设备号',
    key: 'registerMdCode',
    align: 'left',
    width: 150,
    resizable: true,
  },
  {
    title: '注册手机型号',
    key: 'registerMpModel',
    align: 'left',
    width: 120,
  },
  {
    title: '上次登录时间',
    key: 'lastLogin',
    align: 'left',
    width: 170,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 170,
  },
  // {
  //   title: '更新时间',
  //   key: 'updatedAt',
  //   align: 'left',
  //   width: -1,
  // },
];

export const levelList = ref([]);
export const recommendModel = ref('');

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
});

// 加载字典数据选项
export function loadOptions() {
  All({}).then((res) => {
    levelList.value = res.list;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'level':
          item.componentProps.options = levelList.value;
          break;
      }
    }
  });

  new Promise((_resolve, _reject) => {
    getConfig({ group: 'yyconfig' })
      .then((res) => {
        recommendModel.value = res.list.recommendModel;
      })
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
