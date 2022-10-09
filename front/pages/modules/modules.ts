import axios from 'axios'
import { Module, ModuleSummary } from '../../interfaces/modules'


const getModules = async () => {
  try {
    const response = await axios.get(`http://localhost:8080/modules`);
    console.log(response);
    let data = await response.data;
    let modules:Array<ModuleSummary> = data.map( mod => {
      return {
        name: mod.Name,
        providers: mod.Providers
      }
    })
    return modules
  } catch (error) {
    console.log(error)
  }
}

export function ListModules(): Array<ModuleSummary> {
  let modules =  getModules()
  .then( modules => { modules })
  .catch( err => {
    console.log(err)
    return [{name: "", providers: []}]
  })
  console.log("MODULE TYPE: ", typeof modules)
  return []
}