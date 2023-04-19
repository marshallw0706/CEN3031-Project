export class GlobalConstants {
    static get loggedinuser(): string {
        return localStorage.getItem('loggedinuser') || '';
      }
    
      static set loggedinuser(value: string) {
        localStorage.setItem('loggedinuser', value);
      }
    
      static get loggedin(): boolean {
        return localStorage.getItem('loggedin') === 'true';
      }
    
      static set loggedin(value: boolean) {
        localStorage.setItem('loggedin', value.toString());
      }
    
      static get loggedinid(): BigInt {
        return BigInt(localStorage.getItem('loggedinid') || '1');
      }
    
      static set loggedinid(value: BigInt) {
        localStorage.setItem('loggedinid', value.toString());
      }

      static get viewprofileid(): BigInt {
        return BigInt(localStorage.getItem('viewprofileid') || '1');
      }
    
      static set viewprofileid(value: BigInt) {
        localStorage.setItem('viewprofileid', value.toString());
      }

      static get idArray(): number[] {
        const storedIntArray = localStorage.getItem('idArray');
        return storedIntArray ? JSON.parse(storedIntArray) : [];
      }
      
      static set idArray(value: number[]) {
        localStorage.setItem('idArray', JSON.stringify(value));
      }

}