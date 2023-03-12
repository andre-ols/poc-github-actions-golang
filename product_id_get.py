# Library for opening url and creating
# requests
import urllib.request
 
# for parsing all the tables present
# on the website
from html_table_parser import HTMLTableParser
 
import json
from threading import Thread
import time
 

class CustomThread(Thread):
    # constructor
    def __init__(self, vendor_id, arg):
        # execute the base constructor
        Thread.__init__(self, target=self.return_all_product_ids, args=arg)
        # set a default value
        self.value = None
        self.vendor_id = vendor_id
        self.error_num = 1

    def return_all_product_ids(self, vendor_id):
        #try:
        # defining the html contents of a URL.
        vendor_id = vendor_id.upper()
        xhtml = url_get_contents('https://devicehunt.com/search/type/usb/vendor/' + vendor_id +'/device/any').decode('utf-8')
        
        # Defining the HTMLTableParser object
        p = HTMLTableParser()
        
        # feeding the html contents in the
        # HTMLTableParser object
        p.feed(xhtml)
        
        # Now finally obtaining the data of
        # the table required
        if len(p.tables) > 1 :
            products_table = p.tables[1]
            
            products_id_name = {}

            # Ignores first line
            products_table = products_table[1:]
            print(vendor_id)
            for line in products_table:
                product_id = line[3]
                product_name = line[4]

                products_id_name[product_id] = product_name
        
                print(product_name)

            self.value = [vendor_id, products_id_name]
        
        else:
            self.value = {}

        #except Exception as error:
            #print(f"{vendor_id}: Error Number: {self.error_num}" )
            #self.error_num = self.error_num + 1
            #threshold = 0.5
            #time.sleep(threshold)

            #if self.error_num < 10:
            #    self.return_all_product_ids(vendor_id)
            #else:
        #    return

# Opens a website and read its
# binary contents (HTTP Response Body)
def url_get_contents(url):
 
    # Opens a website and read its
    # binary contents (HTTP Response Body)
 
    #making request to the website
    req = urllib.request.Request(url=url)
    f = urllib.request.urlopen(req)
 
    #reading contents of the website
    return f.read()

def getAll():
    with open("vendor_ids2.json") as file:
        vendor_ids = json.load(file)

    product_ids_json = {}

    vendor_keys = list(vendor_ids.keys())

    step_keys = 500

    i = 4

    while True:
        product_ids_json = {}
        return_values = []
        try: 
            #if len(vendor_keys[i*step_keys:]) > (i+1)*step_keys :
            slice_of_keys = vendor_keys[i*step_keys:(i+1)*step_keys]
            #else:
            #    slice_of_keys = vendor_keys[i*step_keys:]
        
            threads = [CustomThread(vendor_id, (vendor_id,)) for vendor_id in slice_of_keys]
            for thread in threads:
                time.sleep(0.01)
                thread.start()

            for thread in threads:
                thread.join()
                return_values.append(thread.value)
            #print(return_values)
            for return_value in return_values:
                if return_value:
                    vendor_id = return_value[0]
                    product_ids_json[vendor_id] = return_value[1]
            
            print(product_ids_json)
            with open(".\product_ids\products_id" + str((i+1)*step_keys) +".json", "w") as file:
                json.dump(product_ids_json, file)

            i = i + 1
        except Exception as error:
            print(error)
            return

if __name__ == '__main__':
    getAll()
    print("Scrapping Complete")
    