<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-grid cols="1" y-gap="15">
        <n-gi>
          <n-card :bordered="false" class="proCard" size="small" :header-style="{
            padding: '25px 20px 15px',
          }" :content-style="{
            padding: '0 20px 22px',
          }">
            <template #header>
              <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">优惠券概况</text>
            </template>
            <n-row>
              <n-col :span="8">
                <div class="stat-div">
                  <div class="stat-div-title">累计发放</div>
                  <div class="stat-div-cont">{{ data.totalSendNum }}</div>
                </div>
              </n-col>
              <n-col :span="8">
                <div class="stat-div">
                  <div class="stat-div-title">已使用</div>
                  <div class="stat-div-cont">{{ data.usedNum }}</div>
                </div>
              </n-col>
              <n-col :span="8">
                <div class="stat-div">
                  <div class="stat-div-title">待使用</div>
                  <div class="stat-div-cont">{{ data.waitUseNum }}</div>
                </div>
              </n-col>
            </n-row>
          </n-card>
        </n-gi>
        <n-gi>
          <n-card :bordered="false" class="proCard" size="small" :header-style="{
            padding: '25px 20px 20px',
          }" :content-style="{
            padding: '0 20px 20px',
          }">
            <template #header>
              <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">优惠券明细</text>
            </template>
            <BasicForm ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable"
              @keyup.enter="reloadTable">
              <template #statusSlot="{ model, field }">
                <n-input v-model:value="model[field]" />
              </template>
            </BasicForm>
            <div class="statusTab w-full hs overflow-x-auto flex-nowrap">
              <div :class="searchState == 0 ? 'active' : ''"><span @click="handleCouponState(0)">全部</span></div>
              <div :class="searchState == 1 ? 'active' : ''"><span @click="handleCouponState(1)">已领取</span></div>
              <div :class="searchState == 2 ? 'active' : ''"><span @click="handleCouponState(2)">已使用</span></div>
              <div :class="searchState == 3 ? 'active' : ''"><span @click="handleCouponState(3)">已过期</span></div>
              <div :class="searchState == 5 ? 'active' : ''"><span @click="handleCouponState(5)">已回收</span></div>
            </div>
            <BasicTable ref="actionRef" :columns="couponColumns2" :request="loadDataTable" :scroll-x="scrollX"
              :resizeHeightOffset="-10000" :actionColumn="actionColumn">
            </BasicTable>
          </n-card>
        </n-gi>
      </n-grid>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref, h, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { NIcon, NTooltip, useMessage, useDialog } from 'naive-ui';
import { Stat, List, Recycle } from '@/api/pmsCoupon';
import { BasicTable, TableAction } from "@/components/Table";
import { adaTableScrollX } from "@/utils/hotgo";
import { BasicForm, useForm } from "@/components/Form";
// import {schemas} from "@/views/pmsBalanceChange/model";
import { couponColumns2, loadCouponOptions, schemas } from "@/views/pmsCouponType/view_coupon_model";
// import {columns} from "@/views/pmsCouponType/model";

const message = useMessage();
const router = useRouter();
const dialog = useDialog();
const params = router.currentRoute.value.params;
const loading = ref(false);
const data = ref({});
const actionRef = ref();
const searchFormRef = ref<any>({});
const [register, { }] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});
const searchState = ref(0);

const actionColumn = reactive({
  width: 90,
  title: '操作',
  key: 'action',
  fixed: 'right',
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '回收',
          onClick: handleRecycle.bind(null, record),
          ifShow: () => {
            return record.state === 1;
          },
          auth: ['/pmsCouponType/delete'],
        },
      ],

    });
  },
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

// 加载表格数据
const loadDataTable = async (res) => {
  if (searchState.value > 0) {
    res.state = searchState.value
  }
  return await List({
    ...{

    }, ...searchFormRef.value?.formModel, ...res
  });
};

const scrollX = computed(() => {
  return adaTableScrollX(couponColumns2, actionColumn.width);
});

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

function handleCouponState(e) {
  searchState.value = e;
  reloadTable()
}

// 回收优惠券
function handleRecycle(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '回收后用户无法再使用，确认回收？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Recycle({ id: record.id }).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          reloadTable();
        });
      });
    },
  });
}

onMounted(() => {
  getInfo();
  loadCouponOptions();
});
</script>

<style lang="less" scoped>
::v-deep(.json-width) {
  width: 100%;
  min-width: 3.125rem;
}

.stat-div {
  padding: 12px 0 19px;

  &-title {
    font-weight: 500;
    font-size: 14px;
    color: #3D3D3D;
    line-height: 20px;
  }

  &-cont {
    margin-top: 5px;
    font-weight: 600;
    font-size: 24px;
    color: #3D3D3D;
    line-height: 34px;
  }
}

.statusTab {
  display: flex;
  margin-bottom: 4px;
  margin-top: 10px;

  div {
    margin-right: 5px;
    padding: 0 16px;
    height: 30px;
    line-height: 30px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    flex-shrink: 0;
    white-space: nowrap;

    span {
      cursor: pointer;
    }

    &.active {
      background: #F2F3F8;
      border-radius: 30px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>
