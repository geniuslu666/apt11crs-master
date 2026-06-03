
// 表格列
import {h,ref} from "vue";
import { Dicts } from '@/api/dict/dict';
import { Option, getOptionLabel, getOptionTag } from '@/utils/hotgo';
import { isNullObject } from '@/utils/is';
import { NTag } from 'naive-ui';

export const couponColumns = [
 /* {
    title: '会员信息',
    key: 'memberId',
    align: 'left',
    width: 200,
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
                width: '50px',
                height: '50px',
                borderRadius: '50%'
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
                {},
                {
                  default: () => row.member.fullName,
                }
              ),
              h(
                'div',
                {},
                {
                  default: () => row.member.phoneArea + ' ' + row.member.phone,
                }
              )
            ]
          )
        ]
      )
    }
  },*/
  {
    title: '优惠券',
    key: 'couponName',
    align: 'left',
    width: 220,
    render(row){
      return row.pmsCouponType.couponName
    }
  },
  {
    title: '类型',
    key: 'type',
    align: 'left',
    width: 90,
    render(row) {
      if(row.type == 'reward'){
        return '满减券'
      }else{
        return '折扣券'
      }
    }
  },
  {
    title: '优惠券金额/折扣',
    key: 'couponName',
    align: 'left',
    width: 130,
    render(record) {
      if(record.type=='reward'){
        return record.money + "JPY";
      }else if(record.type == 'discount'){
        return record.discount + "折";
      }
    }
  },
  {
    title: '满多少元使用',
    key: 'atLeast',
    align: 'left',
    width: 110,
  },
  {
    title: '状态',
    key: 'state',
    align: 'left',
    width: 90,
    render(row) {
      if(row.state == 1){
        return '已领取'
      }else if(row.state == 2){
        return '已使用'
      }else if(row.state == 3){
        return '已过期'
      }else if(row.state == 5){
        return '已回收'
      }else{
        return '已关闭'
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
          'span',
          {
            style: {
              color: 'green',
            },
            bordered: false,
          },
          {
            default: () => '民宿',
          }
        );
      }
      if(record.scene == 2){
        return h(
          'span',
          {
            style: {
              color: 'blue',
            },
            bordered: false,
          },
          {
            default: () => '餐饮',
          }
        );
      }
      if(record.scene == 3){
        return h(
          'span',
          {
            style: {
              color: 'orange',
            },
            bordered: false,
          },
          {
            default: () => '按摩',
          }
        );
      }
      if(record.scene == 4){
        return h(
          'span',
          {
            style: {
              color: 'red',
            },
            bordered: false,
          },
          {
            default: () => '接送机/包车',
          }
        );
      }
      if(record.scene == 5){
        return h(
          'span',
          {
            style: {
              color: 'red',
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
    title: '领取时间',
    key: 'fetchTime',
    align: 'left',
    width: 160,
    ellipsis: false,
  },
  {
    title: '使用时间',
    key: 'useTime',
    align: 'left',
    width: 160,
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

  });
}
