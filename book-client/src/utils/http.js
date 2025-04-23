import axios from 'axios';
import config from '../../appsettings';

const httpClient = axios.create({ baseURL: config.service.baseURL });

export default httpClient;
