<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard"
        size="small"
        :segmented="{ content: true }"
      >
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">积分概况</text>
        </template>
        <n-row>
          <n-col :span="8">
            <n-statistic label="可用积分">
              {{ data.points }}
            </n-statistic>
          </n-col>
          <n-col :span="8">
            <n-statistic label="累计发放积分">
              {{ data.totalSendPoints }}
            </n-statistic>
          </n-col>
          <n-col :span="8">
            <n-statistic label="累计消耗积分">
              {{ data.totalConsumePoints }}
            </n-statistic>
          </n-col>
        </n-row>
      </n-card>
      <n-card
        :bordered="false"
        class="proCard"
        size="small"
        :segmented="{ content: true }"
        style="margin-top: 15px"
      >
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">积分明细</text>
        </template>
        <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
          <template #statusSlot="{ model, field }">
            <n-input v-model:value="model[field]" />
          </template>
          <template #createdAtSlot="{ model, field }">
            <n-date-picker 
              v-model:formatted-value="model[field]" 
              type="datetimerange" 
              value-format="yyyy-MM-dd HH:mm:ss"
              clearable
              :shortcuts="defRangeShortcuts()"
              style="width: 100%"
            />
          </template>
        </BasicForm>
        <BasicTable  ref="actionRef" :columns="columns" :request="loadDataTable" :scroll-x="scrollX" :resizeHeightOffset="-10000">
        </BasicTable>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, h} from 'vue';
import { useRouter } from 'vue-router';
import {NIcon, NTooltip, useMessage} from 'naive-ui';
import { Stat,List } from '@/api/pmsBalanceChange';
import {BasicTable} from "@/components/Table";
import {adaTableScrollX} from "@/utils/hotgo";
import {QuestionCircleOutlined} from "@vicons/antd";
import {BasicForm, useForm} from "@/components/Form";
import { defRangeShortcuts } from '@/utils/dateUtil';
import {schemas} from "@/views/pmsBalanceChange/model";

const message = useMessage();
const router = useRouter();
const params = router.currentRoute.value.params;
const loading = ref(false);
const data = ref({});
const actionRef = ref();
const searchFormRef = ref<any>({});
const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

const getInfo = () => {
  loading.value = true;
  Stat(params)
    .then((res) => {
      data.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
};

// 表格列
const columns = [
  {
    title: '会员',
    key: 'pmsMemberFullName',
    align: 'left',
    width: -1,
    render(row) {
      var mail = row.pmsMemberMail
      if(!mail){
        mail = '无'
      }
      if (row.pmsMemberId <= 0) {
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
              default: () => row.pmsMemberMemberNo
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
              default: () => "编号："+row.pmsMemberMemberNo+"，名称："+row.pmsMemberFullName+"，手机："+row.pmsMemberPhone+"，邮箱：" + mail,
            },
          )
        ],
      );
    },
  },
  {
    title: '方向',
    key: 'changePrice',
    align: 'left',
    width: 110,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: row.changePrice > 0 ? 'green' : 'red'
          },
        },
        {
          default: () => row.changePrice > 0 ? '发放' : '消耗'
        }
      )
    },
  },
  {
    title: '发生场景',
    key: 'scene',
    align: 'left',
    width: 110,
    render(row) {
      if(row.scene == 'SYSTEM'){
        return '系统'
      }else if(row.scene == 'HOTEL'){
        return '酒店'
      }else if(row.scene == 'FOOD'){
        return '餐饮'
      }else if(row.scene == 'SPA'){
        return '按摩'
      }else{
        return '未知'
      }
    },
  },
  {
    title: '描述',
    key: 'des',
    align: 'left',
    width: -1,
  },
  {
    title: '积分变化',
    key: 'changePrice',
    align: 'left',
    width: -1,
    render(row) {
      return h(
        'span',
        {
          style: {
            color: row.changePrice > 0 ? 'green' : 'red'
          },
        },
        {
          default: () => row.changePrice > 0 ? '+' + row.changePrice : row.changePrice
        }
      )
    },
  },
  {
    title: '订单号',
    key: 'orderSn',
    align: 'left',
    width: -1,
  },
  {
    title: '发生时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
];

// 加载表格数据
const loadDataTable = async (res) => {
  return await List({ ...{

    },...searchFormRef.value?.formModel, ...res });
};

const scrollX = computed(() => {
  return adaTableScrollX(columns, 0);
});

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

onMounted(() => {
  getInfo();
});
</script>

<style lang="less" scoped>
::v-deep(.json-width) {
  width: 100%;
  min-width: 3.125rem;
}
</style>


