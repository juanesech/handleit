import axios from 'axios'
import { Module, ModuleSummary } from '../../interfaces/modules'


export const GetModules = async () => {
  let modules:Array<ModuleSummary> = []
  try {
    const response = await axios.get(`http://localhost:8080/modules`);
    let data = await response.data;
    modules = data;
  } catch (err) {
    new Error(`Error listing the modules: ${err}`)
  }
  return modules;
}

export const GetModule = async ( name: string ) => {
  let module: Module = {
    Name: "",
    ID: "",
    Variables: [{
      Name: "",
      Description: "",
      Default: "",
      Required: false,
      Type: ""
    }],
    Outputs: [{
      Name: "",
      Description: ""
    }],
    Providers: [{
      source: "",
      versionConstrains: [""]
    }], 
  };
  try {
    const response = await axios.get(`http://localhost:8080/modules/${name}`);
    console.log(response);
    let data = await response.data;
    module = data;
  } catch (err) {
    new Error(`Error getting module ${name}: ${err}`)
  }
  return module;
}