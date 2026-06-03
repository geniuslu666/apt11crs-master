
// 表格列
import {h} from "vue";
import defaultImg from "@/assets/images/mrtx.png";

export const couponColumns = [
  {
    title: '礼品券编号',
    key: 'couponNo',
    align: 'left',
    width: 155,
  },
  {
    title: '发放时间',
    key: 'createAt',
    align: 'left',
    width: 155,
  },
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
    title: '有效期',
    key: 'endTime',
    align: 'left',
    width: 155,
  },
  {
    title: '状态',
    key: 'state',
    align: 'left',
    width: 90,
    ellipsis: false,
    render(row) {
      if(row.state == 1){
        return '未生效'
      }else if(row.state == 2){
        return '未使用'
      }else if(row.state == 3){
        return '已核销'
      }else if(row.state == 4){
        return '已过期'
      }else if(row.state == 5){
        return '已失效'
      }
    }
  },
  {
    title: '使用时间',
    key: 'verifyTime',
    align: 'left',
    width: 155,
  },
  {
    title: '使用门店',
    key: 'verifyStoreId',
    align: 'left',
    width: 155,
    render(row){
      return row.verifyStore?.storeName
    }
  },
];
