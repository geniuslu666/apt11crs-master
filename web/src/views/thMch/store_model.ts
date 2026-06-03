import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import {NSwitch, NTag, NButton, useMessage} from "naive-ui";
import {Switch} from "@/api/thMchStore";
import {Dicts} from "@/api/dict/dict";
import {isNullObject} from "@/utils/is";
import {Option, getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {Export} from "@/api/thMemberCoupon";

const $message = window['$message'];

export class State {
  public id = 0; // id
  public mchId = null; // 商户ID
  public storeName = ''; // 门店名称
  public nameLanguage = null;
  public images = ''; // 图集
  public imagesArr = [];
  public phoneArea = ''; // 区号
  public phone = ''; // 手机
  public detailAddress = ''; // 详细地址
  public ggLat = '34.67100087743556'; // 谷歌纬度
  public ggLng = '135.49982492658245'; // 谷歌经度
  public terminalIds = ''; // 绑定终端ID
  public terminalList = [];
  public printTerminalIds = ''; // 绑定打印机终端ID
  public handTerminalIds = ''; // 绑定手持终端ID
  public account = ''; // 账号
  public password = ''; // 密码
  public createAt = ''; // create_at
  public updateAt = ''; // update_at
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

// 表格列
export const columns = [
  {
    title: '门店名称',
    key: 'storeName',
    align: 'left',
    width: -1,
  },
  {
    title: '联系信息',
    key: 'phone',
    align: 'left',
    width: -1,
    render(row){
      return h(
        'div',
        {},
        [
          h(
            'div',
            {},
            {
              default: () => row.phoneArea + ' ' + row.phone,
            }
          ),
          h(
            'div',
            {},
            {
              default: () => row.detailAddress,
            }
          )
        ]
      )
    }
  },
  {
    title: '门店状态',
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
    title: '核销数量',
    key: 'verifyNum',
    align: 'left',
    width: -1,
    render(row) {

      return h(
        'div',
        null,
        [
          h(
            'div',
            {
              style: {
                display: 'inline-block',
                minWidth: '40px',
                align: 'center',
              },
            },
            {
              default: () => row.verifyNum,
            }
          ),
          h(
            NButton,
            {
              style:{
                marginLeft: '6px',
              },
              strong: true,
              size: 'small',
              text: true,
              type: 'primary',
              onClick: function (e) {
                $message.loading('正在导出列表...', { duration: 2000 });
                Export({ verifyStoreId: row.id, state: 3 });
              },
            },
            { default: () => '导出' }
          ),
        ]
      )
    },
  },
  {
    title: '是否启用',
    key: 'status',
    align: 'left',
    width: 100,
    render(row) {
      return h(NSwitch, {
        value: row.status === 1,
        checked: '开启',
        unchecked: '关闭',
        onUpdateValue: function (e) {
          row.status = e ? 1 : 2;
          Switch({ id: row.id, key: 'switch', status: row.status }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
];

// 字典数据选项
export const options = ref({
  sys_normal_disable: [] as Option[],
});


export function loadOptions(){
  Dicts({
    types: ['sys_normal_disable'],
  }).then((res) => {
    options.value = res;
  });
}
