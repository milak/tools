package data

type Property struct {
	Name string
	Value interface{}
}
type PropertyList struct {
	propeties []*Property
}
func (this *PropertyList) SetProperty(aName string, aValue interface{}){
	p := this.GetProperty(aName)
	if p == nil {
		this.propeties = append(this.propeties,&Property{Name : aName, Value : aValue})
	} else {
		p.Value = aValue
	}
}
func (this *PropertyList) GetProperty(aName string) *Property {
	for _,p := range this.propeties {
		if p.Name == aName {
			return p
		}
	}
	return nil
}
func (this *PropertyList) GetProperties(){
	return this.propeties
}