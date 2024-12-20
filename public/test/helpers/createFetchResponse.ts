import { FetchResponse } from '@grafana/runtime';

export function createFetchResponse<T>(data: T): FetchResponse<T> {
  return {
    data,
    status: 200,
    url: 'http://localhost:8080/api/tsdb/query',
    config: { url: 'http://localhost:8080/api/tsdb/query' },
    type: 'basic',
    statusText: 'Ok',
    redirected: false,
    headers: ({} as unknown) as Headers,
    ok: true,
  };
}
