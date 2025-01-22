import { useQuery } from '@tanstack/react-query'
import { Link } from "react-router";
import { fetchBatches } from 'api/batches'
import BatchActions from 'components/BatchActions';
import DataTable from 'components/DataTable.tsx'
import { formatDate } from 'util/date';

const Batch = () => {
  const refetchInterval = 10000; // 10 seconds
  const {
    data: batches = [],
    error,
    isLoading,
  } = useQuery({
    queryKey: ['batches'],
    queryFn: fetchBatches,
    refetchInterval: refetchInterval,
  })

  const headers = ["Name", "Status", "Target", "Project", "Storage pool", "Include expression", "Window start", "Window end", "Default network", "Actions"];
  const rows = batches.map((item) => {
    return [
      {
        content: <Link to={`/ui/batches/${item.name}`} className="data-table-link">{item.name}</Link>
      },
      {
        content: item.status_string
      },
      {
        content: item.target_id
      },
      {
        content: item.target_project
      },
      {
        content: item.storage_pool
      },
      {
        content: item.include_expression
      },
      {
        content: formatDate(item.migration_window_start.toString())
      },
      {
        content: formatDate(item.migration_window_end.toString())
      },
      {
        content: item.default_network
      },
      {
        content: <BatchActions batch={item} />
      }];
  });

  if (isLoading) {
    return (
      <div>Loading batches...</div>
    );
  }

  if (error) {
    return (
      <div>Error while loading batches</div>
    );
  }

  return <DataTable headers={headers} rows={rows} />;
};

export default Batch;
