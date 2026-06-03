<template>
  <div>
    <div style="width: 100%;padding-top: 0" class="p-5">
      <div class="mb-10 flex-row" style="margin-bottom: 15px">
        <div class="f30">
          {{ translang('物业中心') }}
          <n-radio-group  name="viewType"  @update:value="handleUpdateViewType" v-model:value="viewType">
            <n-space>
              <n-radio-button :value="1">图片视图</n-radio-button>
              <n-radio-button :value="2">列表视图</n-radio-button>
            </n-space>
          </n-radio-group>
        </div>
<!--        <div class="flex-item text-r">
          <a-input
            v-model:value="actionColumn.searchtext"
            :placeholder="translang('请输入物业名')"
            style="width: 220px"
            @change="searchLoad"
          >
            <template #suffix>
              <a-tooltip title="Extra information">
                <n-icon>
                  <SearchOutline />
                </n-icon>
              </a-tooltip>
            </template>
          </a-input>
           <n-button
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/pmsProperty/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>

          <n-button
            type="primary"
            @click="handleExport"
            class="min-left-space"
            v-if="hasPermission(['/pmsProperty/export'])"
          >
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
        </div>-->
      </div>
    </div>
    <div v-if="viewType==1">
      <div
        style="width: 150px; display: inline-block"
        class="mr-3 mb-10"
        v-for="(item, index) in actionColumn.wylist"
        :key="index"
      >
        <div class="wy bgfff">
          <!-- @click="handleEditimg(item)" -->
          <div class="p-2" style="position: relative">
            <img
              :src="item.cover"
              v-if="item.cover"
              style="
                  top: -15px;
                  width: 135px;
                  height: 170px;
                  object-fit: cover;
                  position: absolute;
                  box-shadow: 0px 0px 4px #d9d9d9;
                "
            />
            <img
              v-else
              style="
                  top: -15px;
                  width: 135px;
                  height: 170px;
                  object-fit: cover;
                  position: absolute;
                  box-shadow: 0px 0px 4px #d9d9d9;
                "
              src="@/assets/images/nowy.png"
            />
            <div style="height: 150px"></div>
            <div class="f16 fw mt-2 nohuanhang">
              <n-tooltip placement="bottom" trigger="hover">
                <template #trigger>
                  <span> {{ namegetlang(item.nameLanguage) }} </span>
                </template>
                <span> {{ namegetlang(item.nameLanguage) }} </span>
              </n-tooltip>
            </div>
            <div class="flex-row">
              <div style="width: 50%" class="c999"
              >{{ translang('房型') }}<span class="ml-2">{{ item.roomTypeNum }}</span>
              </div>
              <div class="c999 flex-item"
              >{{ translang('房间') }}<span class="ml-2">{{ item.roomUnitNum }}</span>
              </div>
            </div>
          </div>
          <div class="flex-row text-c" style="border-top: 1px solid #f7f7f7; cursor: pointer">
            <!--              <div-->
            <!--                style="border-right: 1px solid #f7f7f7; width: 30%"-->
            <!--                class="c333 p-2 f12"-->
            <!--                @click="handleDelete(item)"-->
            <!--                >{{ translang('删除') }}</div-->
            <!--              >-->
            <div
              class="c999 flex-item p-2 f12"
              style="color: #053dc8; border-right: 1px solid #f7f7f7"
              @click="handleEdit(item)"
            >
              <!-- <n-icon size="18" color="#053dc8">
                <CreateOutline />
              </n-icon> -->
              {{ translang('管理') }}
            </div>

            <div
              class="c999 p-2 f12"
              v-if="item.close == 1"
              style="color: #053dc8; width: 50%"
              @click="closepmsProperty(item)"
            >
              {{ translang('已启用') }}
            </div>
            <div
              class="c999 p-2 f12"
              v-else-if="item.close == 2"
              style="color: #999; width: 50%"
              @click="closepmsProperty(item)"
            >
              {{ translang('已关闭') }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- <div class="n-layout-page-header">
      <n-card :bordered="false" title="物业">

      </n-card>
    </div> -->

     <n-card :bordered="false" class="proCard" v-if="viewType==2">
       <div class="statusTab">
         <div :class="tabValue==1 ? 'active' : ''"><span @click="handleUpdateValue('1')">已启用</span></div>
         <div :class="tabValue==2 ? 'active' : ''"><span @click="handleUpdateValue('2')">已关闭</span></div>
         <div :class="tabValue==3 ? 'active' : ''"><span @click="handleUpdateValue('3')">已删除</span></div>
         <div :class="tabValue==0 ? 'active' : ''"><span @click="handleUpdateValue('0')">全部</span></div>
       </div>
      <BasicTable
        ref="actionRef"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :actionColumn="actionColumn"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
        :checked-row-keys="checkedIds"
        :pagination="false"
        @update:checked-row-keys="handleOnCheckedRow"
        @update:sorter="handleUpdateSorter"
      >
        <template #tableTitle>
<!--          <n-button
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/pmsProperty/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button
            type="error"
            @click="handleBatchDelete"
            class="min-left-space"
            v-if="hasPermission(['/pmsProperty/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
          <n-button
            type="primary"
            @click="handleExport"
            class="min-left-space"
            v-if="hasPermission(['/pmsProperty/export'])"
          >
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>-->
        </template>
      </BasicTable>
    </n-card>
    <View ref="viewRef" />
    <EditGroup ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
  import { useDesignSettingStore } from '@/store/modules/designSetting';

  import { h, reactive, ref, computed, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import {List, Export, Delete, Edit, Recycle} from '@/api/pmsProperty';
  // import { pmsmenu } from '@/api/comm';
  import { columns, schemas } from './model';
  import { adaTableScrollX } from '@/utils/hotgo';
  import { getlang, translang } from '@/utils/smjcomm';

  import EditGroup from './edit_group.vue';
  import View from './view.vue';
  import { useUserStore } from '@/store/modules/user';
  import { useSorter } from '@/hooks/common';

  const userStore = useUserStore();

  import { useRouter } from 'vue-router';
  const router = useRouter();

  const dialog = useDialog();

  const designStore = useDesignSettingStore();

  const message = useMessage();
  // const icons =  Ionicons5;
  const compData = ref({
    color: designStore.appTheme,
    size: 32,
  });

  const color = computed(() => {
    return compData.value.color;
  });

  const { hasPermission } = usePermission();
  const actionRef = ref();
  const searchFormRef = ref<any>({});
  const editRef = ref();
  const viewRef = ref();
  const checkedIds = ref([]);

  const viewType = ref(2)
  const tabValue = ref(1)

  const { updateSorter: handleUpdateSorter, sortStatesRef: sortStatesRef } = useSorter(reloadTable);

  const actionColumn = reactive({
    allwylist: [],
    searchtext: '',
    wylist: [],
    pmsdata: [],
    width: 350,
    title: '操作',
    key: 'action',
    fixed: 'right',
    ellipsis: false,
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label:  translang('管理'),
            onClick: handleEdit.bind(null, record),
            auth: ['/pmsProperty/edit'],
          },
          {
            label: translang('关闭'),
            onClick: closepmsProperty.bind(null, record),
            ifShow: () => {
              return record.close === 1;
            },
            auth: ['/pmsProperty/edit'],
          },
          {
            label: translang('启用'),
            onClick: closepmsProperty.bind(null, record),
            ifShow: () => {
              return record.close === 2;
            },
            auth: ['/pmsProperty/edit'],
          },
          {
            label:  translang('绑定'),
            onClick: handleEditGroup.bind(null, record),
            auth: ['/pmsProperty/editGroup'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/pmsProperty/delete'],
            ifShow: () => {
              return record.deletedAt == null;
            },
          },
          {
            label: '恢复',
            onClick: handleRecycle.bind(null, record),
            type: 'warning',
            auth: ['/pmsProperty/delete'],
            ifShow: () => {
              return record.deletedAt != null;
            },
          },
        ],
        dropDownActions: [
          {
            label: '开启预定模式',
            key: 'openBookingClose',
            auth: ['/pmsProperty/edit'],
            ifShow: () => {
              return record.bookingClose === 2;
            },
          },
          {
            label: '关闭预定模式',
            key: 'closeBookingClose',
            auth: ['/pmsProperty/edit'],
            ifShow: () => {
              return record.bookingClose === 1;
            },
          },
          {
            label: '开启短租模式',
            key: 'openLeaseClose',
            auth: ['/pmsProperty/edit'],
            ifShow: () => {
              return record.leaseClose === 2;
            },
          },
          {
            label: '关闭短租模式',
            key: 'closeLeaseClose',
            auth: ['/pmsProperty/edit'],
            ifShow: () => {
              return record.leaseClose === 1;
            },
          },
        ],
        select: (key) => {
          if (key === 'openBookingClose') {
            return handleChangebookingClose(record);
          } else if (key === 'closeBookingClose') {
            return handleChangebookingClose(record);
          } else if (key === 'openLeaseClose') {
            return handleChangeLeaseClose(record);
          } else if (key === 'closeLeaseClose') {
            return handleChangeLeaseClose(record);
          }
        },
      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(columns, actionColumn.width);
  });
  const namegetlang = (data) => {
    if (data) {
      return getlang(data, userStore.language).content;
    } else {
      return '暂无物业名';
    }
  };
  const [register, {}] = useForm({
    /*  */ gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });
  // //PMS物业信息
  // const PmsLoad = async () => {
  //   const res = await pmsmenu({ pid: 2481 });
  //   actionColumn.pmsdata = res.list;
  //   console.log(res);
  // };

  // 加载表格数据
  const loadDataTable = async () => {
    console.log("sortStatesRef", sortStatesRef.value)
    if(sortStatesRef.value.length > 0){
      var sort = "";
      sortStatesRef.value.forEach(element => {
        if(element.columnKey == 'sort'){
          sort = element.order
        }
      });

      var param = { page: 1, pageSize: 200, sort: sort, propertyStatus: tabValue.value }
      if(viewType.value == 2){
        param = { page: 1, pageSize: 200, sort: sort, propertyStatus: tabValue.value }
      }
      const aa = await List(param);
      actionColumn.wylist = aa.list;
      actionColumn.allwylist = aa.list;
      return aa;
    }else{
      var param2 = { page: 1, pageSize: 200, propertyStatus: tabValue.value }
      if(viewType.value == 2){
        param2 = { page: 1, pageSize: 200, propertyStatus: tabValue.value }
      }
      const aa = await List(param2);
      actionColumn.wylist = aa.list;
      actionColumn.allwylist = aa.list;
      return aa;
    }
  };

  // 更新选中的行
  function handleOnCheckedRow(rowKeys) {
    checkedIds.value = rowKeys;
  }

  //
  function searchLoad() {
    actionColumn.wylist = actionColumn.allwylist.filter(
      (data) =>
        !actionColumn.searchtext ||
        data.name.toLowerCase().includes(actionColumn.searchtext.toLowerCase())
    );
  }

  //打开关闭
  function closepmsProperty(record) {
    let content = record.close == 1 ? '确定关闭该物业么？' : '确定开启该物业么？';

    dialog.warning({
      title: '信息',
      content: content,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        const close = record.close == 1 ? 2 : 1;
        Edit({ id: record.id, close: close }).then((_res) => {
          message.success('操作成功');
          if (viewType.value == 1){
            loadDataTable();
          }else{
            reloadTable();
          }
        });
      },
    });
  }

  //打开关闭预定模式
  function handleChangebookingClose(record) {
    let content = record.close == 1 ? '确定关闭预定模式么？' : '确定开启预定模式么？';

    dialog.warning({
      title: '信息',
      content: content,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        const bookingClose = record.bookingClose == 1 ? 2 : 1;

        Edit({
          "bookingClose":bookingClose,
          "id": record.id
        })
          .then((_res) => {
            message.success('操作成功');
            if (viewType.value == 1){
              loadDataTable();
            }else{
              reloadTable();
            }
        });
      },
    });
  }

  //打开关闭短租模式
  function handleChangeLeaseClose(record) {
    let content = record.close == 1 ? '确定关闭短租模式么？' : '确定开启短租模式么？';

    dialog.warning({
      title: '信息',
      content: content,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        const leaseClose = record.leaseClose == 1 ? 2 : 1;

        Edit({
          "leaseClose":leaseClose,
          "id": record.id
        })
          .then((_res) => {
            message.success('操作成功');
            if (viewType.value == 1){
              loadDataTable();
            }else{
              reloadTable();
            }
          });
      },
    });
  }

  function handleUpdateValue(e){
    tabValue.value = e;
    reloadTable()
  }

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  function handleUpdateViewType(value) {
    viewType.value = value;
  }

  // 编辑数据
  function handleEdit(record: Recordable) {
    // editRef.value.openModal(record);
    router.push({ name: 'Propertydetail', query: { id: record.id, uid: record.uid } });
  }

  // 编辑数据
  function handleEditGroup(record: Recordable) {
    editRef.value.openModal(record);
  }

  // 查看详情
  function handleView(record: Recordable) {
    viewRef.value.openModal(record);
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
          if (viewType.value == 1){
            loadDataTable();
          }else{
            reloadTable();
          }
        });
      },
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
          if (viewType.value == 1){
            loadDataTable();
          }else{
            reloadTable();
          }
        });
      },
    });
  }

  // 批量删除
  function handleBatchDelete() {
    if (checkedIds.value.length < 1) {
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
          loadDataTable();
        });
      },
    });
  }

  // 导出
  function handleExport() {
    message.loading('正在导出列表...', { duration: 1200 });
    Export(searchFormRef.value?.formModel);
  }
  onMounted(() => {
    loadDataTable();
  });
</script>

<style lang="less" scoped>
  .meunpms {
    box-shadow: 2px 0 8px 0 rgba(29, 35, 41, 0.05);
    font-size: 14px;
  }
  .pmsitem {
    padding: 12px 12px 12px 22px;
  }
  .pmsitem:hover {
    background-color: hsla(0, 0%, 87%, 0.341);
    border-radius: 5px;
  }

  .statusTab{
    display: flex;
    margin-bottom: 30px;
    div{
      padding: 0 16px;
      height: 28px;
      line-height: 28px;
      text-align: center;
      font-size: 14px;
      color: #4E5969;
      span{
        cursor: pointer;
      }
      &.active{
        background: #F2F3F8;
        border-radius: 28px;
        color: #1664FF;
        font-weight: 500;
      }
    }
  }
</style>
