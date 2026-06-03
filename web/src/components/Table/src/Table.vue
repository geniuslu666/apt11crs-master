<template>
  <div class="table-toolbar">
    <!--顶部左侧区域-->
    <div class="flex items-center table-toolbar-left">
      <template v-if="title">
        <div class="table-toolbar-left-title">
          {{ title }}
          <n-tooltip trigger="hover" v-if="titleTooltip">
            <template #trigger>
              <n-icon size="18" class="ml-1 text-gray-400 cursor-pointer">
                <QuestionCircleOutlined />
              </n-icon>
            </template>
            {{ titleTooltip }}
          </n-tooltip>
        </div>
      </template>
      <slot name="tableTitle"></slot>
    </div>

    <!-- 移动端工具栏右侧区域 -->
    <div class="flex items-center table-toolbar-right-mobile md:hidden">
      <slot name="toolbar"></slot>
    </div>

    <div class="flex items-center table-toolbar-right" v-show="false">
      <!--顶部右侧区域-->
      <slot name="toolbar"></slot>

      <!--斑马纹-->
      <n-tooltip trigger="hover">
        <template #trigger>
          <div class="mr-2 table-toolbar-right-icon">
            <n-switch v-model:value="isStriped" @update:value="setStriped" />
          </div>
        </template>
        <span>表格斑马纹</span>
      </n-tooltip>
      <n-divider vertical />

      <!--刷新-->
      <n-tooltip trigger="hover">
        <template #trigger>
          <div class="table-toolbar-right-icon" @click="reload">
            <n-icon size="18">
              <ReloadOutlined />
            </n-icon>
          </div>
        </template>
        <span>刷新</span>
      </n-tooltip>

      <!--密度-->
      <n-tooltip trigger="hover">
        <template #trigger>
          <div class="table-toolbar-right-icon">
            <n-dropdown @select="densitySelect" trigger="click" :options="densityOptions" v-model:value="tableSize">
              <n-icon size="18">
                <ColumnHeightOutlined />
              </n-icon>
            </n-dropdown>
          </div>
        </template>
        <span>密度</span>
      </n-tooltip>

      <!--表格设置单独抽离成组件-->
      <ColumnSetting :openChecked="openChecked" />
    </div>
  </div>
  <!-- PC端表格 -->
  <div class="s-table hidden md:block">
    <n-data-table ref="tableElRef" v-bind="getBindValues" :striped="isStriped" :pagination="pagination"
      @update:page="updatePage" @update:page-size="updatePageSize">
      <template #[item]="data" v-for="item in Object.keys($slots)" :key="item">
        <slot :name="item" v-bind="data"></slot>
      </template>
    </n-data-table>
  </div>

  <!-- 移动端卡片布局 -->
  <div class="mobile-table-cards block md:hidden">
    <n-spin :show="getLoading">
      <div v-if="tableDataList.length === 0" class="mobile-empty">
        <n-empty description="暂无数据" />
      </div>
      <div v-else class="mobile-cards-container space-y-3">
        <div v-for="(record, index) in tableDataList" :key="getRowKeyValue(record, index)" class="mobile-card">
          <!-- 卡片内容区域 -->
          <div class="mobile-card-content">
            <template v-for="column in displayColumns" :key="column.key">
              <!-- 跳过操作列，操作列单独渲染 -->
              <template v-if="column.key !== 'action'">
                <div class="mobile-card-field">
                  <div class="mobile-card-label">{{ getColumnTitle(column) }}</div>
                  <div class="mobile-card-value">
                    <slot v-if="$slots[column.key]" :name="column.key" :row="record" :column="column" :index="index" />
                    <component v-else-if="column.render" :is="() => column.render(record, index)" />
                    <span v-else>{{ getFieldValue(record, column.key) }}</span>
                  </div>
                </div>
              </template>
            </template>
          </div>

          <!-- 操作列区域 -->
          <div v-if="actionColumnConfig && actionColumnConfig.render" class="mobile-card-actions">
            <div class="mobile-action-wrapper">
              <component :is="() => (actionColumnConfig as any).render(record, index)" />
            </div>
          </div>
        </div>
      </div>

      <!-- 移动端分页 -->
      <div v-if="paginationInfo && typeof paginationInfo === 'object'" class="mobile-pagination mt-4 px-4">
        <n-pagination v-model:page="(paginationInfo as any).page" v-model:page-size="(paginationInfo as any).pageSize"
          :page-count="(paginationInfo as any).pageCount" :item-count="(paginationInfo as any).itemCount"
          :show-size-picker="(paginationInfo as any).showSizePicker !== false"
          :page-sizes="(paginationInfo as any).pageSizes || [10, 20, 30, 50]" @update:page="updatePage"
          @update:page-size="updatePageSize" />
      </div>
    </n-spin>
  </div>
</template>

<script lang="ts">
import {
  ref,
  defineComponent,
  reactive,
  unref,
  toRaw,
  computed,
  toRefs,
  onMounted,
  nextTick,
} from 'vue';
import { ReloadOutlined, ColumnHeightOutlined, QuestionCircleOutlined } from '@vicons/antd';
import { createTableContext } from './hooks/useTableContext';

import ColumnSetting from './components/settings/ColumnSetting.vue';

import { useLoading } from './hooks/useLoading';
import { useColumns } from './hooks/useColumns';
import { useDataSource } from './hooks/useDataSource';
import { usePagination } from './hooks/usePagination';

import { basicProps } from './props';

import { BasicTableProps } from './types/table';

import { getViewportOffset } from '@/utils/domUtils';
import { useWindowSizeFn } from '@/hooks/event/useWindowSizeFn';
import { isBoolean } from '@/utils/is';

const densityOptions = [
  {
    type: 'menu',
    label: '紧凑',
    key: 'small',
  },
  {
    type: 'menu',
    label: '默认',
    key: 'medium',
  },
  {
    type: 'menu',
    label: '宽松',
    key: 'large',
  },
];

export default defineComponent({
  components: {
    ReloadOutlined,
    ColumnHeightOutlined,
    ColumnSetting,
    QuestionCircleOutlined,
  },
  props: {
    ...basicProps,
  },
  emits: [
    'fetch-success',
    'fetch-error',
    'update:checked-row-keys',
    'edit-end',
    'edit-cancel',
    'edit-row-end',
    'edit-change',
  ],
  setup(props, { emit }) {
    const deviceHeight = ref(150);
    const tableElRef = ref<ComponentRef>(null);
    const wrapRef = ref<Nullable<HTMLDivElement>>(null);
    let paginationEl: HTMLElement | null;
    const isStriped = ref(true);
    const tableData = ref<Recordable[]>([]);
    const innerPropsRef = ref<Partial<BasicTableProps>>();

    const getProps = computed(() => {
      return { ...props, ...unref(innerPropsRef) } as BasicTableProps;
    });

    const { getLoading, setLoading } = useLoading(getProps);

    const { getPaginationInfo, setPagination } = usePagination(getProps);

    const { getDataSourceRef, getDataSource, getRowKey, reload } = useDataSource(
      getProps,
      {
        getPaginationInfo,
        setPagination,
        tableData,
        setLoading,
      },
      emit
    );

    const { getPageColumns, setColumns, getColumns, getCacheColumns, setCacheColumnsField } =
      useColumns(getProps);

    // 移动端显示的列（排除操作列）
    const displayColumns = computed(() => {
      return unref(getPageColumns).filter((col) => col.key !== 'action');
    });

    // 获取表格数据列表
    const tableDataList = computed(() => {
      return unref(getDataSourceRef) || [];
    });

    // 获取操作列配置
    const actionColumnConfig = computed(() => {
      const actionCol = unref(getProps).actionColumn;
      if (!actionCol) return null;
      // actionColumn可能是对象，包含render函数
      if (typeof actionCol === 'object' && 'render' in actionCol) {
        return actionCol;
      }
      return null;
    });

    // 获取行key值
    function getRowKeyValue(record, index) {
      const rowKey = unref(getRowKey);
      if (rowKey && typeof rowKey === 'function') {
        return rowKey(record, index);
      }
      return record.id || record.key || index;
    }

    // 获取列标题
    function getColumnTitle(column) {
      if (typeof column.title === 'function') {
        return column.title();
      }
      return column.title || '';
    }

    // 获取字段值
    function getFieldValue(record, key) {
      if (!key) return '--';
      const keys = key.split('.');
      let value = record;
      for (const k of keys) {
        value = value?.[k];
        if (value === undefined || value === null) break;
      }
      return value ?? '--';
    }

    const state = reactive({
      tableSize: unref(getProps as any).size || 'medium',
      isColumnSetting: false,
    });

    //页码切换
    function updatePage(page) {
      setPagination({ page: page });
      reload();
    }

    //分页数量切换
    function updatePageSize(size) {
      setPagination({ page: 1, pageSize: size });
      reload();
    }

    //密度切换
    function densitySelect(e) {
      state.tableSize = e;
    }

    //选中行
    function updateCheckedRowKeys(rowKeys) {
      emit('update:checked-row-keys', rowKeys);
    }

    //获取表格大小
    const getTableSize = computed(() => state.tableSize);

    //组装表格信息
    const getBindValues = computed(() => {
      const tableData = unref(getDataSourceRef);
      const maxHeight = tableData.length ? `${unref(deviceHeight)}px` : 'auto';
      return {
        ...unref(getProps),
        loading: unref(getLoading),
        columns: toRaw(unref(getPageColumns)),
        rowKey: unref(getRowKey),
        data: tableData,
        size: unref(getTableSize),
        remote: true,
        'max-height': maxHeight,
      };
    });

    //获取分页信息
    const pagination = computed(() => toRaw(unref(getPaginationInfo)));
    const paginationInfo = computed(() => toRaw(unref(getPaginationInfo)));

    function setProps(props: Partial<BasicTableProps>) {
      innerPropsRef.value = { ...unref(innerPropsRef), ...props };
    }

    const setStriped = (value: boolean) => (isStriped.value = value);

    const tableAction = {
      reload,
      setColumns,
      setLoading,
      setProps,
      getColumns,
      getPageColumns,
      getCacheColumns,
      setCacheColumnsField,
      emit,
    };

    const getCanResize = computed(() => {
      const { canResize } = unref(getProps);
      return canResize;
    });

    async function computeTableHeight() {
      const table = unref(tableElRef);
      if (!table) return;
      if (!unref(getCanResize)) return;
      const tableEl: any = table?.$el;
      const headEl = tableEl.querySelector('.n-data-table-thead ');
      const { bottomIncludeBody } = getViewportOffset(headEl);
      const headerH = 64;
      let paginationH = 2;
      let marginH = 24;
      if (!isBoolean(unref(pagination))) {
        paginationEl = tableEl.querySelector('.n-data-table__pagination') as HTMLElement;
        if (paginationEl) {
          const offsetHeight = paginationEl.offsetHeight;
          paginationH += offsetHeight || 0;
        } else {
          paginationH += 28;
        }
      }
      let height =
        bottomIncludeBody - (headerH + paginationH + marginH + (props.resizeHeightOffset || 0));
      const maxHeight = props.maxHeight;
      height = maxHeight && maxHeight < height ? maxHeight : height;
      deviceHeight.value = height;
    }

    useWindowSizeFn(computeTableHeight, 280);

    onMounted(() => {
      nextTick(() => {
        computeTableHeight();
      });
    });

    createTableContext({ ...tableAction, wrapRef, getBindValues });

    return {
      ...toRefs(state),
      tableElRef,
      getBindValues,
      getDataSource,
      densityOptions,
      reload,
      densitySelect,
      updatePage,
      updatePageSize,
      pagination,
      tableAction,
      setStriped,
      isStriped,
      displayColumns,
      tableDataList,
      actionColumnConfig,
      getColumnTitle,
      getFieldValue,
      getRowKeyValue,
      getLoading,
      getRowKey,
      getDataSourceRef,
      paginationInfo,
    };
  },
});
</script>
<style lang="less" scoped>
.table-toolbar {
  display: flex;
  justify-content: space-between;
  padding: 0 0 16px 0;

  &-left {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    flex: 1;

    &-title {
      display: flex;
      align-items: center;
      justify-content: flex-start;
      font-size: 16px;
      font-weight: 600;
    }
  }

  &-right {
    display: flex;
    justify-content: flex-end;
    flex: 1;

    &-icon {
      margin-left: 12px;
      font-size: 16px;
      cursor: pointer;
      color: var(--text-color);

      :hover {
        color: #1890ff;
      }
    }
  }
}

.table-toolbar-inner-popover-title {
  padding: 2px 0;
}

/* 移动端卡片样式 */
.mobile-table-cards {
  width: 100%;
}

.mobile-empty {
  padding: 40px 16px;
  text-align: center;
}

.mobile-cards-container {
  padding: 0;
}

.mobile-card {
  transition: all 0.2s ease;
  background: #ffffff;
  border-radius: 8px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
  border: 1px solid #f0f0f0;
  overflow: hidden;
}

.mobile-card:active {
  transform: scale(0.98);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.mobile-card-content {
  padding: 12px 16px;
}

.mobile-card-field {
  display: flex;
  align-items: center;
  min-height: 32px;
  margin-bottom: 8px;
}

.mobile-card-field:last-child {
  margin-bottom: 0;
}

.mobile-card-label {
  font-weight: 500;
  font-size: 13px;
  line-height: 20px;
  color: #6b7280;
  flex-shrink: 0;
  margin-right: 12px;
}

.mobile-card-value {
  font-size: 14px;
  line-height: 20px;
  color: #111827;
  word-break: break-word;
  flex: 1;
  text-align: right;
}

.mobile-card-actions {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  border-top: 1px solid #f0f0f0;
  padding: 12px 16px;
  background: #f9fafb;
}

.mobile-action-wrapper {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 6px;
  flex-wrap: wrap;
  width: 100%;
}

/* 移动端操作按钮样式优化 - 移除固定高度，让size属性生效 */
.mobile-card-actions :deep(.n-button) {
  border-radius: 6px;
}

/* 确保small size的按钮在移动端正确显示 */
.mobile-card-actions :deep(.n-button--small) {
  height: 32px !important;
  min-height: 32px !important;
  padding: 0 12px !important;
  font-size: 13px !important;
  line-height: 32px !important;
}

.mobile-card-actions :deep(.tableAction) {
  width: 100%;
}

.mobile-card-actions :deep(.tableAction .flex) {
  justify-content: flex-end;
  flex-wrap: wrap;
  gap: 8px;
}

/* 移动端分页样式优化 */
.mobile-pagination {
  padding: 16px;
  display: flex;
  justify-content: center;
}

.mobile-pagination :deep(.n-pagination) {
  justify-content: center;
}

.mobile-pagination :deep(.n-pagination-item) {
  min-width: 40px;
  min-height: 40px;
}

.mobile-pagination {
  display: flex;
  justify-content: center;
}

/* 移动端工具栏适配 */
.table-toolbar {
  @media (max-width: 767px) {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
    padding: 0 0 12px 0;

    &-left {
      width: 100%;

      &-title {
        font-size: 16px;
        font-weight: 600;
      }
    }

    &-right-mobile {
      width: 100%;
      justify-content: flex-start;
      flex-wrap: wrap;
      gap: 8px;

      :deep(.n-button) {
        min-height: 36px;
        font-size: 13px;
      }
    }
  }
}
</style>
