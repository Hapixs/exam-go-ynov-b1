# Exam GO

* ### Fonction **IsAlpha(r rune)bool** <br>
    * **Description** <br> La fonction à pour but de retourner si oui ou non une rune 'r' est une lettre.

    * **Fonctionnement** <br> La fonction regarde et compare la valeurs ASCII du la rune 'r'. <br> Si elle correspond à une valeur de lettre (de 65 à 90 ou 97 à 122) alors la fonction retourne ```true``` sinon elle retourne ```false```

    * **Example** <br> 
    ```go 
    fmt.Println(exam.IsAlpha('a'))
	fmt.Println(exam.IsAlpha('R'))
	fmt.Println(exam.IsAlpha('2'))
    ```
    ```go
    true  // 'a' est une lettre
    true  // 'R' est une lettre
    false // '2' n'est pas une lettre
    ```

* ### Fonction **Average(t []int)int**
    * **Description** <br> La fonction à pour but de calculer la moyenne des valeurs contenue dans un tableau de nombre 't'

    * **Fonctionnement** <br> La fonction vas boucler sur tous les nombres du tableau en les ajoutant les uns au autres. <br> Elle va ensuite diviser le total par le nombre de fois ou elle à été modifier. <br> Dans le cas ou le tableau contient ```-1```, cette valeur ne sera pas prise en compte dans l'ajout ni dans la division.

    * **Example**
    ```go
    fmt.Println(exam.Average([]int{14, 12, 7, 5, 17}))
	fmt.Println(exam.Average([]int{13, -1, 8, 15, 15}))
    ``` 
    ```go
    11 // Ici, 14+12+7+5+17 = 55 (total); On divise le total par le nombre de nombres positifs contenue dans le tableau donc 55/5=11

    12 // Ici, 13+8+15+15 = 51 (total); On divise le total par le nombre de nombres positifs contenue dans le tableau donc 51/4=12
    ```

* ### Fonction **PrintElements(tab []int, str []string)string**
    * **Description** <br> La fonction retourne un string contenant tout les valeurs du tableau 'str' correspondante au index contenue dans 'tab'

    * **Fonctionnement** <br> La fonction vas boucler sur 'tab' et gardant la valeur dans 'value' <br> 
    Elle vas ensuite ajouter au string final la valeur de str à l'index 'value' de notre boucle <br>
    
    * **Example**
    ```go
    fmt.Println(exam.PrintElements([]int{1, 4}, []string{"Choice", "Hello", "Solid", "Curtain", "There", "Forward"}))
	fmt.Println(exam.PrintElements([]int{2, 0, 7}, []string{"Kenobi", "Unity", "General", "Therapist"}))
    ```
    ```go
    Hello There
    General Kenobi !
    ``` 
    ![alt text](https://i.kym-cdn.com/photos/images/newsfeed/001/475/420/c62.gif)

* ### Programme **use-arguments**
    * **Description** <br>

    * **Fonctionnement** <br>

    * **Example** <br>
    ```go
    ```
    ```go
    ```

* ### Fonction **RecursiveSequence(i, max int)int**
    * **Description** <br>

    * **Fonctionnement** <br>

    * **Example** <br>
    ```go
    ```
    ```go
    ```

* ### Fonction **Decipher(s string)string**
    * **Description** <br>

    * **Fonctionnement** <br>

    * **Example** <br>
    ```go
    ```
    ```go
    ```
* ### Fonction **Diamond(index int)**
    * **Description** <br>

    * **Fonctionnement** <br>

    * **Example** <br>
    ```go
    ```
    ```go
    ```
* ### Fonction **Conway(index int)int**
    * **Description** <br>

    * **Fonctionnement** <br>

    * **Example** <br>
    ```go
    ```
    ```go
    ```

![alt text](https://thumbs.gfycat.com/DazzlingIckyAtlanticbluetang-max-1mb.gif)