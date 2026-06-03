import {h, ref} from 'vue';
import { FormSchema } from '@/components/Form';
import {defRangeShortcuts} from '@/utils/dateUtil';
import defaultImg from "@/assets/images/mrtx.png";
import {isNullObject} from "@/utils/is";
import {NTag} from "naive-ui";
import {getOptionLabel, getOptionTag, Option} from "@/utils/hotgo";
import {Dicts} from "@/api/dict/dict";

// 表格搜索表单
export const juniorSchemas = ref<FormSchema[]>([
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
    field: 'createdAt',
    component: 'NDatePicker',
    label: '发生时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const juniorColumns = [
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
                  default: () => row.fullName,
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
      )
    },
  },

  {
    title: '会员积分/经验',
    key: 'exp',
    align: 'left',
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
];

export const options = ref({
  sys_normal_disable: [] as Option[],
});

// 加载字典数据选项
export function loadJuniorOptions() {
  Dicts({
    types: ['sys_normal_disable'],
  }).then((res) => {
    options.value = res;
  });
}
