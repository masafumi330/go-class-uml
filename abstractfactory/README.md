# Abstract Factory

- Abstract Factoryパターンは、**生成に関するパターン**である。
- **生成に関するパターン**は、柔軟性を増し、コードの再利用を促すようなコードの仕組みを提供する。
- Abstract Factoryパターンは、関連したオブジェクトの集まりを、具象クラスを指定することなく生成できる。
- つまり、クライアント側は、具体的なクラスを指定する必要はなく、「なんとなくこれが欲しい」というふわっとした粒度でオブジェクトの生成を依頼することができる。具体的に「どう」生成するかは、Abstract Factoryパターンの実装のなかで決められるということだ。


```mermaid
---
title: 常に同じブランドの靴とシャツを取得するAbstract Factory
--- 

classDiagram
    class Usecase {
        -factory ISportsFactory
    }

    Usecase ..|> ISportsFactory : "nikeFactory = factory.GetSportsFactory('nike')"

    class ISportsFactory {
        <<interface>>
        makeShoe() IShoe
        makeShirt() IShirt
    }

    namespace 具象ファクトリー {
        class Adidas {
            makeShoe() IShoe
            makeShirt() IShirt
        }
        class Nike {
            makeShoe() IShoe
            makeShirt() IShirt
        }
    }

    ISportsFactory <|.. Adidas
    ISportsFactory <|.. Nike

    namespace 具象プロダクト {
        class AdidasShoe {
            
        }
        class NikeShoe {
        }
        class AdidasShirt {
            
        }
        class NikeShirt {
        }
    }

    namespace 抽象プロダクト {
        class IShoe {
            <<interface>>
            -setLogo(string)
            -setSize(int)
            -getLogo() string
            -getSize() int
        }
        class IShirt {
            <<interface>>
            -setLogo(string)
            -setSize(int)
            -getLogo() string
            -getSize() int
        }
    }

    Nike ..> NikeShoe
    Nike ..> NikeShirt
    Adidas ..> AdidasShoe
    Adidas ..> AdidasShirt

    class Shoe {
        -logo string
        -size int
        -setLogo(string)
        -setSize(int)
        -getLogo() string
        -getSize() int
    }
    class Shirt {
        -logo string
        -size int
        -setLogo(string)
        -setSize(int)
        -getLogo() string
        -getSize() int
    }

    NikeShoe *.. Shoe
    AdidasShoe *.. Shoe
    NikeShirt *.. Shirt
    AdidasShirt *.. Shirt

    Shirt ..|> IShirt
    Shoe ..|> IShoe
```