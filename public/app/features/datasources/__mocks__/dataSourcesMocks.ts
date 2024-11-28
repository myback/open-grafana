import { DataSourceSettings } from '@grafana/data';

// kludges

export const getMockDataSources = (amount: number) => {
  const dataSources = [];

  for (let i = 0; i < amount; i++) {
    dataSources.push({
      access: '',
      basicAuth: false,
      database: `database-${i}`,
      id: i,
      isDefault: false,
      jsonData: { authType: 'credentials', defaultRegion: 'eu-west-2' },
      name: `dataSource-${i}`,
      orgId: 1,
      password: '',
      readOnly: false,
      type: 'elasticsearch',
      typeLogoUrl: 'public/app/plugins/datasource/elasticsearch/img/elasticsearch.svg',
      url: '',
      user: '',
    });
  }

  return dataSources as DataSourceSettings[];
};

export const getMockDataSource = (): DataSourceSettings => {
  return {
    access: '',
    basicAuth: false,
    basicAuthUser: '',
    basicAuthPassword: '',
    withCredentials: false,
    database: '',
    id: 13,
    isDefault: false,
    jsonData: { authType: 'credentials', defaultRegion: 'eu-west-2' },
    name: 'gdev-cloudwatch',
    typeName: 'Elasticsearch',
    orgId: 1,
    password: '',
    readOnly: false,
    type: 'elasticsearch',
    typeLogoUrl: 'public/app/plugins/datasource/elasticsearch/img/elasticsearch.svg',
    url: '',
    user: '',
    secureJsonFields: {},
  };
};
