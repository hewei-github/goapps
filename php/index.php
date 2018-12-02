<?php

    // echo time() . PHP_EOL;
    // echo phpversion() . PHP_EOL;
    if(!empty($_GET)){
        foreach($_GET as $key =>$val){
            echo "${key} : ${val} </br>"  ;
        }
    }
    echo "</br>";
    echo phpinfo() . PHP_EOL;