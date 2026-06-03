<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">服务列表</text>
        </template>
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard" :content-style="{
                    padding: '0 20px 20px',
                  }">
<!--      <BasicForm  ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">-->
<!--        <template #statusSlot="{ model, field }">-->
<!--          <n-input v-model:value="model[field]" />-->
<!--        </template>-->
<!--      </BasicForm>-->
      <n-tabs type="line" animated v-model:value="tabValue" @update:value="handleUpdateValue">
        <n-tab-pane :name="0" tab="全部">
        </n-tab-pane>
        <n-tab-pane :name="item.value" :tab="item.label" v-for="item in options.service_status">
        </n-tab-pane>
        <n-tab-pane :name="3" tab="回收站">
        </n-tab-pane>
      </n-tabs>
      <BasicTable  ref="actionRef" :singleLine="false" :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
        <template #tableTitle>
          <n-button type="primary"  @click="addTable" class="min-left-space" v-if="hasPermission(['/spaService/edit'])">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
<!--          <n-button type="error" @click="handleBatchDelete" class="min-left-space" v-if="hasPermission(['/spaService/delete'])">-->
<!--            <template #icon>-->
<!--              <n-icon>-->
<!--                <DeleteOutlined />-->
<!--              </n-icon>-->
<!--            </template>-->
<!--            批量删除-->
<!--          </n-button>-->
        </template>
      </BasicTable>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, onMounted } from 'vue';
import {NButton, NImage, NTag, useDialog, useMessage} from 'naive-ui';
import { BasicTable, TableAction } from '@/components/Table';
import { BasicForm, useForm } from '@/components/Form/index';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Status, Delete, Recycle } from '@/api/spaService';
import { PlusOutlined } from '@vicons/antd';
import { schemas, options, loadOptions } from './model';
import {adaTableScrollX, getOptionLabel, getOptionTag} from '@/utils/hotgo';
import {useRouter} from "vue-router";
import {isNullObject} from "@/utils/is";

const router = useRouter();
const tabValue = ref(0)
const dialog = useDialog();
const message = useMessage();
const { hasPermission } = usePermission();
const actionRef = ref();
const searchFormRef = ref<any>({});
const checkedIds = ref([]);

const columns = [
  {
    title: '缩略图',
    key: 'id',
    align: 'left',
    width: 120,
    rowSpan: (rowData, rowIndex) => rowData.priceList.length,
    render(row){
      if(row.images){
        let imagesArr = row.images.split(',');
        return h(
          NImage,
          {
            width: 100,
            src: imagesArr[0],
          },

        )
      }else{
        return '--'
      }
    }
  },
  {
    title: '服务名称',
    key: 'name',
    align: 'left',
    width: 200,
    rowSpan: (rowData, rowIndex) => rowData.priceList.length,
  },
  {
    title: '所属服务商',
    key: 'ispId',
    align: 'left',
    width: 120,
    rowSpan: (rowData, rowIndex) => rowData.priceList.length,
    render(row){
      if(row.ispId > 0){
        return row.spaIspDetail.name
      }else{
        return '--'
      }
    }
  },
  {
    title: '服务标签',
    key: 'labelIds',
    align: 'left',
    width: 250,
    rowSpan: (rowData, rowIndex) => rowData.priceList.length,
    render(row){
      var labelList = [];
      if(row.labelList){
        row.labelList.forEach((item) => {
          var labelH = h(
            NTag,
            {
              style: {
                marginRight: '6px',
              },
              type: 'primary',
              bordered: false,
              size: 'small'
            },
            {
              default: () => item.labelDetail.labelName,
            }
          )
          labelList.push(labelH)
        })
      }else{
        var labelH = h(
          'div',
          null,
          {
            default: () => '--',
          }
        )
        labelList.push(labelH)
      }

      return h(
        'div',
        null,
        labelList
      )
    }
  },
  {
    title: '项目名称',
    key: 'goodsName',
    align: 'left',
    width: 200,
    render(row){
      var labelH = ""

      if(row.goodsStatus == 2){
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
            default: () => "禁用",
          }
        )
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
              default: () => row.goodsName,
            }
          ),
          labelH
        ]

      )
    }
  },
  {
    title: '项目价格',
    key: 'goodsPrice',
    align: 'left',
    width: 100,
    render(row){
      return row.goodsPrice + 'JPY'
    }
  },
  {
    title: '项目时长',
    key: 'goodsDuration',
    align: 'left',
    width: 100,
    render(row){
      return row.goodsDuration + '分钟'
    }
  },
  {
    title: '排序',
    key: 'sort',
    align: 'left',
    width: 80,
    rowSpan: (rowData, rowIndex) => rowData.priceList.length,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 110,
    rowSpan: (rowData, rowIndex) => rowData.priceList.length,
    render(row) {
      if (isNullObject(row.serviceState)) {
        return ``;
      }
      return h(
        NButton,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.service_status, row.serviceState),
          onClick: () => {
            Status({ id: row.id, serviceStatus: row.serviceState == 1 ? 2 : 1 }).then((_res) => {
              message.success('设为' + getOptionLabel(options.value.service_status,  row.serviceState == 1 ? 2 : 1) + '成功');
              setTimeout(() => {
                reloadTable();
              });
            });
          }
        },
        {
          default: () => getOptionLabel(options.value.service_status, row.serviceState),
        }
      );
    },
  },
  {
    title: '时间',
    key: 'createAt',
    align: 'left',
    width: 250,
    rowSpan: (rowData, rowIndex) => rowData.priceList.length,
    render(row){
      return h(
        'div',
        null,
        [
          h(
            'div',
            null,
            {
              default: () => '创建时间：' + row.createAt,
            }
          ),
          h(
            'div',
            null,
            {
              default: () => '更新时间：' + row.updateAt,
            }
          )
        ]
      )
    }
  },
];

const actionColumn = reactive({
  width: 288,
  title: '操作',
  key: 'action',
  fixed: 'right',
  rowSpan: (rowData, rowIndex) => rowData.priceList.length,
  render(record) {
    return h(TableAction as any, {
      style: 'button',
      actions: [
        {
          label: '编辑',
          onClick: handleEdit.bind(null, record),
          ifShow: () => {
            return record.deletedAt  == null;
          },
          auth: ['/spaService/edit'],
        },
        {
          label: '复制',
          onClick: handleCopy.bind(null, record),
          ifShow: () => {
            return record.deletedAt  == null;
          },
          auth: ['/spaService/edit'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          ifShow: () => {
            return record.deletedAt  == null;
          },
          auth: ['/spaService/delete'],
        },
        {
          label: '恢复',
          onClick: handleRecycle.bind(null, record),
          type: 'warning',
          auth: ['/spaService/recycle'],
          ifShow: () => {
            return record.deletedAt != null;
          },
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(columns, actionColumn.width);
});

const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

function handleUpdateValue(e){
  tabValue.value = e;
  reloadTable()
}

// 加载表格数据
const loadDataTable = async (res) => {
  if(tabValue.value != 0){
    res.serviceStatus = tabValue.value
  }
  let dataPageList = await List({ ...searchFormRef.value?.formModel, ...res });
  let dataList = dataPageList.list;
  let newDataList = [];
  dataList.forEach((item) => {
    if(item.priceList.length > 0){
      item.priceList.forEach((item1) => {
        let priceItem = {
          'goodsName': item1.goodsName,
          'goodsPrice': item1.price,
          'goodsDuration': item1.duration,
          'goodsStatus': item1.status
        }
        newDataList.push({...item,...priceItem})
      })
    }
  });
  dataPageList.list = newDataList;
  return dataPageList
  // return await List({ ...searchFormRef.value?.formModel, ...res });
};

// 更新选中的行
function handleOnCheckedRow(rowKeys) {
  checkedIds.value = rowKeys;
}

// 重新加载表格数据
function reloadTable() {
  actionRef.value?.reload();
}

// 添加数据
function addTable() {
  router.push({ name: 'spaServiceEdit', params: { id: 0 } });
}

// 编辑数据
function handleEdit(record: Recordable) {
  router.push({ name: 'spaServiceEdit', params: { id: record.id } });
}

// 复制数据
function handleCopy(record: Recordable) {
  router.push({ name: 'spaServiceCopy', params: { id: record.id } });
}

// 单个删除
function handleDelete(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '你确定要删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete(record).then((_res) => {
        message.success('删除成功');
        reloadTable();
      });
    },
  });
}

// 批量删除
function handleBatchDelete() {
  if (checkedIds.value.length < 1){
    message.error('请至少选择一项要删除的数据');
    return;
  }

  dialog.warning({
    title: '警告',
    content: '你确定要批量删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete({ id: checkedIds.value }).then((_res) => {
        checkedIds.value = [];
        message.success('删除成功');
        reloadTable();
      });
    },
  });
}

// 修改状态
function handleStatus(record: Recordable, status: number) {
  Status({ id: record.id, status: status }).then((_res) => {
    message.success('设为' + getOptionLabel(options.value.service_status, status) + '成功');
    setTimeout(() => {
      reloadTable();
    });
  });
}

// 单个恢复
function handleRecycle(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '你确定要恢复？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Recycle(record).then((_res) => {
        message.success('恢复成功');
        reloadTable();
      });
    },
  });
}

onMounted(() => {
  loadOptions();
});
</script>

<style lang="less" scoped></style>

