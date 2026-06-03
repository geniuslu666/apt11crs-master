<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">员工管理详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              基本信息
            </div>
            <n-grid y-gap="25" :cols="3">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">员工姓名</div>
                  <div class="level-detail-div-item-content">
                    {{ data.name }}
                    <a href="javascript:void(0);" @click="handleEditBase('name', data)">修改</a>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">员工部门</div>
                  <div class="level-detail-div-item-content">
                    {{ data.department }}
                    <a href="javascript:void(0);" @click="handleEditBase('department', data)">修改</a>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">手机号</div>
                  <div class="level-detail-div-item-content">
                    {{ data.phone }}
                    <a href="javascript:void(0);" @click="handleEditBase('phone', data)">修改</a>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">邮箱</div>
                  <div class="level-detail-div-item-content">
                    {{ data.email }}
                    <a href="javascript:void(0);" @click="handleEditBase('email', data)">修改</a>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">返佣比例</div>
                  <div class="level-detail-div-item-content">
                    {{ data.rate }}%
                    <a href="javascript:void(0);" @click="handleEditBase('rate', data)">修改</a>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">最低可提现额</div>
                  <div class="level-detail-div-item-content">
                    {{ data.minWithdrawalAmount }}JPY
                    <a href="javascript:void(0);" @click="handleEditBase('minWithdrawalAmount', data)">修改</a>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">提现手续费</div>
                  <div class="level-detail-div-item-content">
                    {{ data.serviceCharge }}%
                    <a href="javascript:void(0);" @click="handleEditBase('serviceCharge', data)">修改</a>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">提现周期</div>
                  <div class="level-detail-div-item-content">
                    {{ data.afterDay }}天
                    <a href="javascript:void(0);" @click="handleEditBase('afterDay', data)">修改</a>
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">绑定会员</div>
                  <div class="level-detail-div-item-content">
                    {{ data.pmsMemberId > 0 ? data.pmsMemberFullName : '--' }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">状态</div>
                  <div class="level-detail-div-item-content">
                    <div class="level-detail-div-item-content-tag1 success" v-if="data.status == 1">正常</div>
                    <div class="level-detail-div-item-content-tag1 warning" v-else>停用</div>
                    <a href="javascript:void(0);" @click="handleEditBase('status', data)">修改</a>
                  </div>
                </div>
              </n-gi>
            </n-grid>
            <div class="level-detail-div-title" style="margin-top: 25px">
              <div></div>
              账户信息
            </div>
            <n-grid y-gap="25" :cols="3">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">总余额</div>
                  <div class="level-detail-div-item-content">
                    {{ data.allBalance }}JPY
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">可提现余额</div>
                  <div class="level-detail-div-item-content">
                    {{ data.balance }}JPY
                  </div>
                </div>
              </n-gi>
            </n-grid>
            <div class="level-detail-div-title" style="margin-top: 25px;margin-bottom: 0">
              <div></div>
              结算订单
            </div>
            <BasicTable  ref="actionRef" :columns="columns" :request="loadDataTable" :scroll-x="scrollX" :resizeHeightOffset="-10000">
            </BasicTable>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

    <EditBase ref="editBaseRef" @reloadInfo="reloadInfo" />
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, h} from 'vue';
import { View } from '@/api/pmsStaff';
import { rebateList } from '@/api/pmsAppReservation';
import {BasicTable} from "@/components/Table";
import {adaTableScrollX, adaModalWidth} from "@/utils/hotgo";
import EditBase from "@/views/pmsStaff/edit_base.vue";

const emit = defineEmits(['reloadTable']);
const showModal = ref(false);
const loading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});
const staffId = ref(0)
// 字典数据选项
const data = ref({});
const editBaseRef = ref();

const getInfo = async () => {
  loading.value = true;
  const res =  await View({id: staffId.value})
  data.value = res;
  loading.value = false;
};

// 表格列
const columns = [
  {
    title: '订单号',
    key: 'orderSn',
    align: 'left',
    width: -1,
  },
  {
    title: '订单金额',
    key: 'orderAmount',
    align: 'left',
    width: -1,
    render: function (rowData){
      return h(
        'span',
        {
          bordered: false
        },
        [
          h(
            'span',
            {
              bordered: false
            },
            {
              default: () => rowData.orderAmount
            }
          ),
          h(
            'span',
            {
              class: 'c999',
              style: {
                marginLeft: '6px',
              },
            },
            {
              default: () => 'JPY'
            }
          )
        ]
      )
    }
  },
  {
    title: '退款金额',
    key: 'refundAmount',
    align: 'left',
    width: -1,
  },
  {
    title: '分佣比例',
    key: 'rebateRate',
    align: 'left',
    width: -1,
    render: function (rowData){
      return rowData.rebateRate + '%'
    }
  },
  {
    title: '结算金额',
    key: 'rebateAmount',
    align: 'left',
    width: -1,
  },
  {
    title: '结算状态',
    key: 'rebateStatus',
    align: 'left',
    width: -1,
    render: function (rowData){
      if(rowData.rebateStatus == 'WAIT'){
        return h(
          'div',
          {
            style: {
              color: '#919399',
              padding: '0px 5px',
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
            default: () => '结算中',
          }
        );
      }else if(rowData.rebateStatus == 'SUCCESS'){
        return h(
          'div',
          {
            style: {
              color: '#26A763',
              padding: '0px 5px',
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
            default: () => '结算成功',
          }
        );
      }else{
        return h(
          'div',
          {
            style: {
              color: '#F56C6C',
              padding: '0px 5px',
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
            default: () => '结算失败',
          }
        );
      }
    }
  },
];

// 加载表格数据
const loadDataTable = async (res) => {
  return await rebateList({ ...{
      fx_type: 'staff',
      fx_id: staffId.value
    }, ...res });
};

const scrollX = computed(() => {
  return adaTableScrollX(columns, 0);
});

// 编辑数据
function handleEditBase(type, value) {
  editBaseRef.value.openModal(data.value.id, type, value);
}

async function reloadInfo(){
  await getInfo();
  emit('reloadTable');
}

async function openModal(id) {
  staffId.value = id
  showModal.value = true;
  await getInfo();
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped>
::v-deep(.json-width) {
  width: 100%;
  min-width: 3.125rem;
}

.level-detail-div{
  margin-bottom: 25px;
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 25px;
    div{
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }
  &-item{
    &-title{
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      margin-bottom: 8px;
    }
    &-content{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      display: flex;
      align-items: center;
      a{
        font-weight: 400;
        font-size: 14px;
        color: #1577FF;
        line-height: 22px;
        margin-left: 6px;
        cursor: pointer;
      }
      .color{
        width: 20px;
        height: 20px;
        margin-right: 8px;
        border-radius: 2px;
      }
      &-tag1{
        padding: 0 5px;
        height: 22px;
        line-height: 22px;
        text-align: center;
        font-size: 14px;
        border-radius: 2px;
        font-weight: 400;
        &.success{
          color: #26A763;
          background: #E3F4EB;
        }
        &.warning{
          color: #EFA020;
          background: #FFECCE;
        }
        &.primary{
          color: #3F9EFF;
          background: #ECF5FF;
        }
        &.error{
          color: #F56C6C;
          background: #FEF0F0;
        }
      }
    }
  }
}
</style>


