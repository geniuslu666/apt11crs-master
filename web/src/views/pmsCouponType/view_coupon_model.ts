
// 表格列
import {h, ref} from "vue";
import { FormSchema } from '@/components/Form';
import defaultImg from "@/assets/images/mrtx.png";
import {isNullObject} from "@/utils/is";
import {NTag} from "naive-ui";
import { Dicts } from '@/api/dict/dict';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'memberKey',
    component: 'NInput',
    label: '会员信息',
    componentProps: {
      placeholder: '请输入会员名/手机号/邮箱',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'scene',
    component: 'NSelect',
    label: '适用场景',
    componentProps: {
      placeholder: '请选择',
      options: [
        {
          labelField: '民宿',
          valueField: 1,
        },
        {
          labelField: '订餐',
          valueField: 2,
        },
        {
          labelField: '按摩',
          valueField: 3,
        },
        {
          labelField: '接送机/包车',
          valueField: 4,
        },
        {
          labelField: '储物柜',
          valueField: 5,
        },
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
 /* {
    field: 'source',
    component: 'NSelect',
    label: '获取方式',
    componentProps: {
      placeholder: '请选择',
      options: [
        {
          labelField: '自主领取',
          valueField: '1',
        },
        {
          labelField: '系统发放',
          valueField: '2',
        },
        {
          labelField: '注册奖励',
          valueField: '3',
        },
        {
          labelField: '邀请奖励',
          valueField: '4',
        },
      ],
      labelField: 'labelField',
      valueField: 'valueField',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },*/
  {
    field: 'source',
    component: 'NSelect',
    label: '获取方式',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },

]);

export const couponColumns = [
  {
    title: '会员信息',
    key: 'memberId',
    align: 'left',
    width: 250,
    render(row) {
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
              }
            }
          ),
          h(
            'div',
            {
              class:'flex-item',
              style: {
                paddingLeft: '8px',
              }
            },
            [
              h(
                'div',
                {
                  style: {
                    fontWeight: '400',
                    fontSize: '14px',
                    color: '#1664FF',
                    lineHeight: '20px'
                  }
                },
                {
                  default: () => row.member.memberNo,
                }
              ),
              h(
                'div',
                {
                  style: {
                    fontWeight: '400',
                    fontSize: '14px',
                    color: '#3D3D3D',
                    lineHeight: '20px'
                  }
                },
                {
                  default: () => row.member.fullName + ' ' + row.member.phoneArea + ' ' + row.member.phone,
                }
              ),
            ]
          )
        ]
      )
    }
  },
  // {
  //   title: '优惠券',
  //   key: 'couponName',
  //   align: 'left',
  //   width: 180,
  //   ellipsis: false,
  //   render(row){
  //     return row.pmsCouponType.couponName
  //   }
  // },
  // {
  //   title: '类型',
  //   key: 'type',
  //   align: 'left',
  //   width: 100,
  //   render(row) {
  //     if(row.type == 'reward'){
  //       return '满减券'
  //     }else{
  //       return '折扣券'
  //     }
  //   }
  // },
  // {
  //   title: '适用场景',
  //   key: 'scene',
  //   align: 'left',
  //   width: 90,
  //   render(record) {
  //     if(record.scene == 1){
  //       return h(
  //         'span',
  //         {
  //           style: {
  //             color: 'green',
  //           },
  //           bordered: false,
  //         },
  //         {
  //           default: () => '民宿',
  //         }
  //       );
  //     }else if(record.scene == 2){
  //       return h(
  //         'span',
  //         {
  //           style: {
  //             color: 'blue',
  //           },
  //           bordered: false,
  //         },
  //         {
  //           default: () => '订餐',
  //         }
  //       );
  //     }else if(record.scene == 3){
  //       return h(
  //         'span',
  //         {
  //           style: {
  //             color: 'orange',
  //           },
  //           bordered: false,
  //         },
  //         {
  //           default: () => '按摩',
  //         }
  //       );
  //     }else if(record.scene == 4){
  //       return h(
  //         'span',
  //         {
  //           style: {
  //             color: 'red',
  //           },
  //           bordered: false,
  //         },
  //         {
  //           default: () => '接送机/包车',
  //         }
  //       );
  //     }
  //     return '--'
  //   }
  // },
  {
    title: '领取来源',
    key: 'source',
    align: 'left',
    width: 100,
    render(row) {
      if (isNullObject(row.source)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.coupon_source, row.source),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.coupon_source, row.source),
        }
      );
    },
  },
  {
    title: '状态',
    key: 'state',
    align: 'left',
    width: 90,
    ellipsis: false,
    render(row) {
      if(row.state == 1){
        return '已领取'
      }else if(row.state == 2){
        return '已使用'
      }else if(row.state == 3){
        return '已过期'
      }else if(row.state == 4){
        return '已关闭'
      }else if(row.state == 5){
        // return '已回收'
        return h(
          'div',
          {
            style: {
              // display: 'flex',
              // alignItems: 'center'
              lineHeight: "15px"
            }
          },
          [
            h(
              'p',
              {
                style: {
                  color: "red"
                }
              },
              {
                default: () => "已回收",
              }
            ),
            h(
              'p',
              {},
              {
                default: () => row.recoveryTime,
              }
            ),
          ]
        )
      }
    }
  },
  {
    title: '领取时间',
    key: 'fetchTime',
    align: 'left',
    width: 155,
    ellipsis: false,
  },
  {
    title: '使用时间',
    key: 'useTime',
    align: 'left',
    width: 155,
  },
];

export const couponColumns2 = [
  {
    title: '会员信息',
    key: 'memberId',
    align: 'left',
    width: 250,
    render(row) {
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
              }
            }
          ),
          h(
            'div',
            {
              class:'flex-item',
              style: {
                paddingLeft: '8px',
              }
            },
            [
              h(
                'div',
                {
                  style: {
                    fontWeight: '400',
                    fontSize: '14px',
                    color: '#1664FF',
                    lineHeight: '20px'
                  }
                },
                {
                  default: () => row.member.memberNo,
                }
              ),
              h(
                'div',
                {
                  style: {
                    fontWeight: '400',
                    fontSize: '14px',
                    color: '#3D3D3D',
                    lineHeight: '20px'
                  }
                },
                {
                  default: () => row.member.fullName + ' ' + row.member.phoneArea + ' ' + row.member.phone,
                }
              ),
            ]
          )
        ]
      )
    }
  },
  {
    title: '优惠券',
    key: 'couponName',
    align: 'left',
    width: 180,
    ellipsis: false,
    render(row){
      return row.pmsCouponType.couponName
    }
  },
  {
    title: '类型',
    key: 'type',
    align: 'left',
    width: 80,
    render(row) {
      if(row.type == 'reward'){
        return h(
          'div',
          {
            style: {
              color: '#17A158',
              width: '52px',
              height: '22px',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
              border: '1px solid #17A158'
            },
            bordered: false,
          },
          {
            default: () => '满减券',
          }
        );
      }
      if(row.type == 'discount'){
        return h(
          'div',
          {
            style: {
              color: '#F48720',
              width: '52px',
              height: '22px',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
              border: '1px solid #F48720'
            },
            bordered: false,
          },
          {
            default: () => '折扣券',
          }
        );
      }
    }
  },
  {
    title: '适用场景',
    key: 'scene',
    align: 'left',
    width: 90,
    render(record) {
      if(record.scene == 1){
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
            default: () => '民宿',
          }
        );
      }else if(record.scene == 2){
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
            default: () => '订餐',
          }
        );
      }else if(record.scene == 3){
        return h(
          'div',
          {
            style: {
              color: '#EFA020',
              padding: '0 5px',
              height: '22px',
              background: '#FFECCE',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '按摩',
          }
        );
      }else if(record.scene == 4){
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
            default: () => '接送机/包车',
          }
        );
      }else if(record.scene == 5){
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
            default: () => '储物柜',
          }
        );
      }
      return '--'
    }
  },
  {
    title: '领取来源',
    key: 'source',
    align: 'left',
    width: 90,
    render(row) {
      if (isNullObject(row.source)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.coupon_source, row.source),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.coupon_source, row.source),
        }
      );
    },
  },
  {
    title: '状态',
    key: 'state',
    align: 'left',
    width: 150,
    ellipsis: false,
    render(row) {
      if(row.state == 1){
        return h(
          'div',
          {
            style: {
              color: '#26A763',
              width: '52px',
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
            default: () => '已领取',
          }
        );
      }
      if(row.state == 2){
        return h(
          'div',
          {
            style: {
              color: '#919399',
              width: '52px',
              height: '22px',
              background: '#F4F4F5',
              lineHeight: '22px',
              textAlign: 'center',
              fontSize: '14px',
              borderRadius: '2px',
              fontWeight: '400',
            },
            bordered: false,
          },
          {
            default: () => '已使用',
          }
        );
      }
      if(row.state == 3){
        return h(
          'div',
          {},
          [
            h(
              'div',
              {
                style: {
                  color: '#919399',
                  width: '52px',
                  height: '22px',
                  background: '#F4F4F5',
                  lineHeight: '22px',
                  textAlign: 'center',
                  fontSize: '14px',
                  borderRadius: '2px',
                  fontWeight: '400',
                },
                bordered: false,
              },
              {
                default: () => '已过期',
              }
            ),
            h(
              'div',
              {
                style: {
                  color: '#3D3D3D',
                  lineHeight: '20px',
                  fontSize: '14px',
                  marginTop: '3px',
                  fontWeight: '400',
                },
                bordered: false,
              },
              {
                default: () => row.endTime,
              }
            )
          ]
        );
      }
      if(row.state == 4){
        return h(
          'div',
          {
            style: {
              color: '#F56C6C',
              width: '52px',
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
            default: () => '已关闭',
          }
        );
      }
      if(row.state == 5){
        return h(
          'div',
          {},
          [
            h(
              'div',
              {
                style: {
                  color: '#F56C6C',
                  width: '52px',
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
                default: () => '已回收',
              }
            ),
            h(
              'div',
              {
                style: {
                  color: '#3D3D3D',
                  lineHeight: '20px',
                  fontSize: '14px',
                  marginTop: '3px',
                  fontWeight: '400',
                },
                bordered: false,
              },
              {
                default: () => row.recoveryTime,
              }
            )
          ]
        );
      }
    }
  },
  {
    title: '领取时间',
    key: 'fetchTime',
    align: 'left',
    width: 155,
    ellipsis: false,
  },
  {
    title: '使用时间',
    key: 'useTime',
    align: 'left',
    width: 155,
  },
];

// 字典数据选项
export const options = ref({
  coupon_source: [] as Option[],
});


// 加载字典数据选项
export function loadCouponOptions() {
  Dicts({
    types: ['coupon_source'],
  }).then((res) => {
    options.value = res;
    for (const item of schemas.value) {
      switch (item.field) {
        case 'source':
          item.componentProps.options = options.value.coupon_source;
          break;
      }
    }
  });
}
