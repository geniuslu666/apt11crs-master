import { h, ref } from 'vue';
import {NTag, NInput, NSwitch} from 'naive-ui';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import {getlang} from "@/utils/smjcomm";
import { useUserStore } from '@/store/modules/user';
import {UpdateAccessPass, SwitchSpaCanOrder} from "@/api/pmsProperty";

const userStore = useUserStore();

const $message = window['$message'];


// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'uid',
    component: 'NInput',
    label: '物业ID',
    componentProps: {
      placeholder: '请输入物业ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
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
export const columns = [
  {
    title: '物业名称',
    key: 'name',
    align: 'left',
    width: 300,
    render: function (row){
      var labelH = ""

      if(row.deletedAt != null){
        labelH = h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'error',
            bordered: false,
            size: 'small'
          },
          {
            default: () => "已删除",
          }
        )
      }
      var propertyName = "--"
      if(row.nameLanguage){
        propertyName = getlang(row.nameLanguage, userStore.language).content
      }

      return h(
        'div',
        null,
        [
          h(
            'span',
            {
              style: {
                display: "inline-block",
                marginRight: '6px'
              }
            },
            {
              default: () => propertyName,
            }
          ),
          labelH
        ]

      )
    }
  },
  {
    title: '房型数',
    key: 'roomTypeNum',
    align: 'left',
    width: 80,
  },
  {
    title: '房间数',
    key: 'roomUnitNum',
    align: 'left',
    width: 80,
  },
  // {
  //   title: '地址',
  //   key: 'address',
  //   align: 'left',
  //   width: -1,
  //   render: function(row){
  //     if (row.addressLanguage) {
  //       return getlang(row.addressLanguage, userStore.language).content;
  //     } else {
  //       return '--';
  //     }
  //   }
  // },
  {
    title: '排序',
    key: 'sort',
    sorter: true, // 单列排序
    width: 80,
    render(row) {
      return row.sort
      // return h(NInput, {
      //   value: row.sort,
      //   onUpdateValue(v) {
      //     Sort({ id: row.id, sort: v }).then((_res) => {
      //       $message.success('操作成功');
      //     });
      //     row.sort = v
      //   }
      // })
    }
  },
  {
    title: '状态',
    key: 'name',
    align: 'left',
    width: 80,
    render: function(row){
      if (row.close ==1){
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: "success",
            bordered: false,
          },
          {
            default: () => "已启用",
          }
        );
      }else{
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: "warning",
            bordered: false,
          },
          {
            default: () => "已关闭",
          }
        );
      }

    }
  },
  {
    title: '是否开放预约',
    key: 'spaCanOrder',
    align: 'left',
    width: 80,
    render(row) {
      return h(NSwitch, {
        value: row.spaCanOrder === 1,
        checked: '开放预约',
        unchecked: '关闭预约',
        onUpdateValue: function (e) {
          row.spaCanOrder = e ? 1 : 2;
          SwitchSpaCanOrder({ id: row.id, key: 'switch', spaCanOrder: row.spaCanOrder }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
  {
    title: '门禁密码',
    key: 'accessPass',
    width: 100,
    render(row) {
      return h(NInput, {
        value: row.accessPass,
        onUpdateValue(v) {
          row.accessPass = v
        },
        onBlur() {
          UpdateAccessPass({ id: row.id, accessPass: row.accessPass }).then((_res) => {
            $message.success('操作成功');
          });
        },
      })
    }
  },
];
